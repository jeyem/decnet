package decnet

import (
	"bytes"
	"fmt"
	"io"
	"net"
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
	DBDir       string
}

type Connection struct {
	listner  int
	db       *badger.DB
	handlers map[string]HandlerFunc
}

func New(opt Options, key *Key) (*Connection, error) {
	conn := new(Connection)
	conn.handlers = map[string]HandlerFunc{}
	conn.listner = opt.Port
	db, err := badger.Open(badger.DefaultOptions(opt.DBDir))
	if err != nil {
		return nil, err
	}
	conn.db = db
	return conn, nil
}

func (c *Connection) AddHandler(action string, handler HandlerFunc) {
	c.handlers[action] = handler
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

	// txn := c.db.NewTransaction(true)

	var (
		buf    = make([]byte, 8000)
		first  = true
		action string
	)

	context := c.newContext()
	context.tcpConnection = conn
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		packet := new(Packet)
		if err := proto.Unmarshal(buf[:n], packet); err != nil {
			c.replayString(conn, "REQUEST NOT ACCEPTTED", RejectAction)
			return
		}

		packet.Action == TouchAction

		context.request.Write(packet.Body)
		h, ok := c.handlers[packet.Action]
		if !ok {
			c.replayString(conn, "COULD NOT FIND ACTION HANDLER", RejectAction)
			break
		}
		if first {
			if err := h(context); err != nil {
				logrus.Error()
				return
			}
			action = packet.Action
		}
		first = false
		if packet.Completed {
			break
		}
	}
	c.replay(conn, context.response, action)
	logrus.Warn("socket closed")

}

func (c *Connection) replayString(conn net.Conn, message, action string) error {
	reader := bytes.NewReader([]byte(message))
	return c.replay(conn, reader, action)
}

func (c *Connection) replay(conn net.Conn, body io.Reader, action string) error {
	var buf = make([]byte, 6000)
	for {
		n, err := body.Read(buf)
		if err != nil {
			break
		}
		packet, err := proto.Marshal(&Packet{
			Action:    action + "_replay",
			Listener:  int32(c.listner),
			Username:  "TODO",
			PublicKey: "TODO",
			Headers:   "TODO",
			Completed: n < len(buf),
			Body:      buf[:n],
			Created:   time.Now().Unix(),
		})

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
	for {
		n, err := body.Read(buf)
		if err != nil {
			break
		}
		packet, err := proto.Marshal(&Packet{
			Action:    action,
			Listener:  int32(c.listner),
			Username:  "TODO",
			PublicKey: "TODO",
			Headers:   "TODO",
			Completed: n < len(buf),
			Body:      buf[:n],
			Created:   time.Now().Unix(),
		})

		if err != nil {
			return err
		}

		if _, err := conn.Write(packet); err != nil {
			return err
		}
	}

	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		p := new(Packet)
		if err := proto.Unmarshal(buf[:n], p); err != nil {
			return err
		}

		logrus.Info(p)
	}

	return nil
}
