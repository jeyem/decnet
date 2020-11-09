package decnet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

/* func TestDB(t *testing.T) {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	txn := db.NewTransaction(true)
	c := new(compeer)
	c.Listener = "127.0.0.1:8080"
	c.Username = "Test"

	if err := c.save(txn); err != nil {
		t.Error(err)
		return
	}
	if err := txn.Commit(); err != nil {
		t.Error(err)
		return
	}
	txn2 := db.NewTransaction(true)

	com, err := findPeer(c.Listener, txn2)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("set username -> ", c.Username, "   read username  -> ", com.Username)
}
*/
func TestEcho(t *testing.T) {
	const (
		APort = 8000
		BPort = 8001
	)
	keyA, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	keyB, err := GenerateKey()
	t.Error(err)
	if err != nil {
		return
	}
	connA, err := New(Options{Port: APort, DBPath: "decnet_test_conn_a.db"}, keyA)
	if err != nil {
		t.Error(err)
		return
	}
	connA.AddHandler("echo", func(c *Context) error {
		return c.Replay("Echo")
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
		return c.Replay(string(data))
	})

	go connA.Start()
	go connB.Start()

	time.Sleep(time.Second)

	reader := bytes.NewReader([]byte("hi"))
	if err := connA.Send(fmt.Sprintf("0.0.0.0:%d", BPort), reader, "echo"); err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 5)

}
