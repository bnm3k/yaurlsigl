package store

import (
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

// NewStore initiates and returns an instance of store. It also
// ensures that the necessary buckets are set up.
// An application should only have one instance of Store.
func NewStore(db *bolt.DB, infoLog, errorLog *log.Logger) (*Store, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return errors.Wrapf(err, "create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Store{
		db:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
	}, nil
}
