package decnet

import (
	"crypto/rsa"
	"encoding/json"
	"time"

	"github.com/dgraph-io/badger"
)

type compeer struct {
	ID        string    `json:"id"`
	PublicKey string    `json:"public_key`
	Listener  string    `json:"listener"`
	Updated   time.Time `json:"updated"`
}

func (c *compeer) save(txn *badger.Txn) error {
	c.Updated = time.Now()
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return txn.Set([]byte(c.Listener), data)
}

func (c *compeer) getPublicKey() *rsa.PublicKey {
	k, _ := convertBytesToPublicKey([]byte(c.PublicKey))
	return k
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
