package decnet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestEncryption(t *testing.T) {
	const message = "this is secret message !"
	key, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	encrypted, err := Encrypt(key.publicKey, []byte(message))
	if err != nil {
		t.Error(err)
		return
	}
	result, err := key.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
		return
	}
	if string(result) != message {
		t.Error(string(result), "not matched with ", message)
	}
	logrus.Info(string(result))
}

func TestEcho(t *testing.T) {
	const (
		message = "Hello Sir!"
		APort   = 8000
		BPort   = 8001
	)
	keyA, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	keyB, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	connA, err := New(Options{Port: APort, DBPath: "decnet_test_conn_a.db"}, keyA)
	if err != nil {
		t.Error(err)
		return
	}
	connA.AddReplayHandler("echo", func(c *Context) error {
		data, err := ioutil.ReadAll(c.Body())
		if err != nil {
			return err
		}
		if string(data) != message {
			t.Errorf("%s not match with > %s", string(data), message)
		}
		return nil
	})
	connB, err := New(Options{Port: BPort, DBPath: "decnet_test_conn_b.db"}, keyB)
	if err != nil {
		t.Error(err)
		return
	}
	connB.AddHandler("echo", func(c *Context) error {
		data, err := ioutil.ReadAll(c.Body())
		if err != nil {
			return err
		}
		if string(data) != message {
			t.Errorf("%s not match with > %s", string(data), message)
		}
		return c.Replay(string(data))
	})

	go connA.Start()
	go connB.Start()

	time.Sleep(time.Second)

	reader := bytes.NewReader([]byte(message))
	if err := connA.Send(fmt.Sprintf("0.0.0.0:%d", BPort), reader, "echo"); err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second)
}
