package decnet

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

const (
	TouchAction  = "Touch"
	RejectAction = "Rejected"
)

type HandlerFunc func(c *Context) error

type Options struct {
	Port        int
	StartPoints []string
	DBPath      string
}

type Connection struct {
	listner  int
	db       *badger.DB
	handlers map[string]HandlerFunc
	key      *Key
}

func New(opt Options, key *Key) (*Connection, error) {
	conn := new(Connection)
	conn.handlers = map[string]HandlerFunc{}
	conn.listner = opt.Port
	conn.key = key
	db, err := badger.Open(badger.DefaultOptions(opt.DBPath))
	if err != nil {
		return nil, err
	}
	conn.db = db
	return conn, nil
}

func (c *Connection) AddHandler(action string, handler HandlerFunc) {
	if action == TouchAction || action == RejectAction {
		logrus.Fatalf("could not use %s action", action)
	}
	c.handlers[action] = handler
}

func (c *Connection) AddReplayHandler(action string, handler HandlerFunc) {
	c.AddHandler(action+"_replay", handler)
}

func (c *Connection) Start() {
	c.tcpListener()
}

func (c *Connection) tcpListener() {
	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", c.listner))
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go c.handler(conn)
	}
}

func (c *Connection) handler(conn net.Conn) {
	defer conn.Close()

	txn := c.db.NewTransaction(true)
	defer txn.Commit()

	var (
		buf    = make([]byte, 8000)
		first  = true
		action string
	)
	peer := new(compeer)
	context := c.newContext()
	context.tcpConnection = conn
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		packet := new(Packet)
		if err := proto.Unmarshal(buf[:n], packet); err != nil {
			c.sendString(conn, "REQUEST NOT ACCEPTTED", RejectAction, nil)
			return
		}
		if packet.Action == TouchAction {
			peer.ID = packet.Sender
			peer.Listener = fmt.Sprintf("%s:%d", strings.Split(conn.RemoteAddr().String(), ":")[0], packet.Listener)
			peer.PublicKey = string(packet.Body)
			peer.save(txn)
			c.sendString(conn, c.key.PublicKeyToPemString(), TouchAction, nil)
			continue
		}
		body, err := c.key.Decrypt(packet.Body)
		if err != nil {
			logrus.Error(err)
			return
		}
		context.request.Write(body)
		h, ok := c.handlers[packet.Action]
		if !ok {
			c.sendString(conn, "COULD NOT FIND ACTION HANDLER", RejectAction, nil)
			break
		}
		if first {
			if err := h(context); err != nil {
				logrus.Error(err)
				return
			}
			action = packet.Action
		}
		first = false
		if packet.Completed {
			break
		}
	}
	c.send(conn, context.response, action+"_replay", peer.getPublicKey())
}

func (c *Connection) sendString(conn net.Conn, message, action string, publicKey *rsa.PublicKey) error {
	reader := bytes.NewReader([]byte(message))
	return c.send(conn, reader, action, publicKey)
}

func (c *Connection) send(conn net.Conn, body io.Reader, action string, publicKey *rsa.PublicKey) error {
	var buf = make([]byte, 6000)
	for {
		n, err := body.Read(buf)
		if err != nil {
			break
		}
		data, _ := Encrypt(publicKey, buf[:n])
		packet, err := c.makePacket(action, data, n < len(buf))
		if err != nil {
			return err
		}

		if _, err := conn.Write(packet); err != nil {
			return err
		}
	}
	return nil
}

func (c *Connection) Send(listener string, body io.Reader, action string) error {
	var buf = make([]byte, 6000)
	conn, err := net.Dial("tcp", listener)
	if err != nil {
		return err
	}

	if err := c.sendString(conn, c.key.PublicKeyToPemString(), TouchAction, nil); err != nil {
		return err
	}

	txn := c.db.NewTransaction(true)
	defer txn.Commit()

	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	packet := new(Packet)
	if err := proto.Unmarshal(buf[:n], packet); err != nil {
		c.sendString(conn, "REQUEST NOT ACCEPTTED", RejectAction, nil)
		return err
	}

	peer := new(compeer)
	peer.ID = packet.Sender
	peer.Listener = fmt.Sprintf("%s:%d", strings.Split(conn.RemoteAddr().String(), ":")[0], packet.Listener)
	peer.PublicKey = string(packet.Body)
	peer.save(txn)

	if err := c.send(conn, body, action, peer.getPublicKey()); err != nil {
		return err
	}

	first := true
	context := c.newContext()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		p := new(Packet)
		if err := proto.Unmarshal(buf[:n], p); err != nil {
			return err
		}
		body, err := c.key.Decrypt(p.Body)
		if err != nil {
			logrus.Error(err)
			continue
		}
		context.request.Write(body)
		h, ok := c.handlers[p.Action]
		if !ok {
			break
		}
		if first {
			if err := h(context); err != nil {
				return err
			}
			action = packet.Action
		}
		first = false
	}
	return nil
}

func (c *Connection) makePacket(action string, content []byte, completed bool, headers ...string) ([]byte, error) {
	return proto.Marshal(&Packet{
		Action:    action,
		Listener:  int32(c.listner),
		Headers:   headers,
		Completed: completed,
		Body:      content,
		Created:   time.Now().Unix(),
	})
}
