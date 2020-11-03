package decnet

import (
	"bytes"
	"fmt"
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
	connA, _ := New(Options{Port: APort})
	connA.AddHandler("echo", func(c *Context) error {
		return c.Replay("Echo")
	})

	connB, _ := New(Options{Port: BPort})
	connB.AddHandler("echo", func(c *Context) error {
		return c.Replay("Echo")
	})

	go connA.Start()
	go connB.Start()

	time.Sleep(time.Second * 5)

	reader := bytes.NewReader([]byte("hi"))
	if err := connA.Send(fmt.Sprintf("0.0.0.0:%d", BPort), reader, "echo"); err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 20)

}
