package store

import (
	"fmt"
	"net/url"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

// ErrInvalidURL is returned when the URL provided by a user
// for shortening is not a valid URL eg /foo/bar
var ErrInvalidURL = errors.New("URL provided is invalid")

// ErrInvalidShortcode is returned when the shortcode provided by a user
// does not exist in the db
var ErrInvalidShortcode = errors.New("URL Shortcode entry does not exist")

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

// ShortenURL given a url, validates the URL, stores it then
// returns a shortcode which can be used as a shortened version
// of the URL
func (s *Store) ShortenURL(fullURL string) (shortened string, err error) {
	//validate URL
	_, err = url.ParseRequestURI(fullURL)
	if err != nil {
		return "", ErrInvalidURL
	}
	//generate shortcode
	shortcode, err := shortid.Generate()
	if err != nil {
		return "", err
	}
	//insert into db
	err = s.db.Update(func(tx *bolt.Tx) error {
		bktName := []byte("urls")
		b := tx.Bucket(bktName)
		if b == nil {
			return fmt.Errorf("faulty internal store setup. Bucket with given name %s not present", bktName)
		}
		if err != nil {
			return errors.Wrap(err, "Creating bucket failed")
		}
		return b.Put([]byte(shortcode), []byte(fullURL))
	})

	return shortcode, err
}

//GetFullURL returns the full URL given the short code
func (s *Store) GetFullURL(shortened string) (fullURL string, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		bktName := []byte("urls")
		b := tx.Bucket(bktName)
		if b == nil {
			return fmt.Errorf("faulty internal store setup. Bucket with given name %s not present", bktName)
		}
		val := b.Get([]byte(shortened))
		if val == nil {
			return ErrInvalidShortcode
		}
		fullURL = string(val)
		return nil
	})
	return
}
