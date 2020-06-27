package store

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// Store encapsulates all CRUD methods for domain
type Store struct {
	db *bolt.DB
}

// NewStore initiates and returns an instance of store. It also
// ensures that the necessary buckets are set up.
// An application should only have one instance of Store.
func NewStore(dbPath string) (*Store, error) {
	// db
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open db: %s", dbPath)
	}

	err = db.Update(func(tx *bolt.Tx) error {
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
		db: db,
	}, nil
}

// Close closes the store. Should be called only once
// after one's done
func (s *Store) Close() error {
	return s.db.Close()
}
