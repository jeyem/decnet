package decnet

import (
	"encoding/json"
	"time"

	"github.com/dgraph-io/badger"
)

type compeer struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Listener string    `json:"listener"`
	Updated  time.Time `json:"updated"`
}

func (c *compeer) save(txn *badger.Txn) error {
	c.Updated = time.Now()
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return txn.Set([]byte(c.Listener), data)
}

func findPeer(key string, txn *badger.Txn) (*compeer, error) {
	item, err := txn.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	var data []byte
	result, err := item.ValueCopy(data)
	if err != nil {
		return nil, err
	}
	c := new(compeer)
	if err := json.Unmarshal(result, c); err != nil {
		return nil, err
	}
	return c, nil
}
