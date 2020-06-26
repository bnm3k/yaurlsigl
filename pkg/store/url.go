package store

import (
	"errors"
	"net/url"

	"github.com/teris-io/shortid"
)

// ErrInvalidURL is returned when the URL provided by a user
// for shortening is not a valid URL eg /foo/bar
var ErrInvalidURL = errors.New("URL provided is invalid")

// ErrInvalidShortcode is returned when the shortcode provided by a user
// does not exist in the db
var ErrInvalidShortcode = errors.New("URL Shortcode entry does not exist")

// ErrUnableToShortenURL error occurs if shortid fails
// Note: shortID fails if configured digits/chars are invalid
var ErrUnableToShortenURL = errors.New("Unable to shorten URL")

// ErrDuplicateURL occurs when user tries to generate shortcode for the same
// url more than once, it's recommended that user deactives/deletes previous entry
// if they wish to insert a new
var ErrDuplicateURL = errors.New("url already shortened for given user")

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
		return "", ErrUnableToShortenURL
	}
	//insert into db

	// return
	return shortcode, nil
}

//GetFullURL returns the full URL given the short code
func (s *Store) GetFullURL(shortened string) (fullURL string, err error) {
	return
}
