package store

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Store encapsulates all CRUD methods for domain
type Store struct {
	db       *bolt.DB
	infoLog  *log.Logger
	errorLog *log.Logger
}

// NewStore initiates and returns an instance of store,
// An application should only have one instance of Store
func NewStore(db *bolt.DB, infoLog, errorLog *log.Logger) *Store {
	return &Store{
		db:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
	}
}

func put(db *bolt.DB, key, val []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("foo"))
		if err != nil {
			return errors.Wrap(err, "Creating bucket failed")
		}
		err = b.Put(key, val)
		return err
	})
	return err
}

func get(db *bolt.DB, key []byte) error {
	var val []byte
	err := db.View(func(tx *bolt.Tx) error {
		bktName := []byte("foo")
		b := tx.Bucket(bktName)
		if b == nil {
			return fmt.Errorf("bucket with given name %s not present", bktName)
		}
		val = b.Get(key)
		if val == nil {
			return fmt.Errorf("entry with given key %s not present", key)
		}
		//fmt.Printf("%s -> %s\n", key, val)
		return nil
	})
	return err
}
