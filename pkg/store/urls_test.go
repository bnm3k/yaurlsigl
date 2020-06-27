package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortenThenRetrieveURL(t *testing.T) {
	// temp dir
	dirPath, err := ioutil.TempDir(".", "test")
	if err != nil {
		os.RemoveAll(dirPath)
		t.Fatal(err)
	}

	// db
	testDBPath := filepath.Join(dirPath, "test.db")
	store, err := NewStore(testDBPath)
	defer func() {
		store.Close()
		os.RemoveAll(dirPath)
	}()
	require.NoError(t, err, "NewStore should not fail when args are valid")

	// shorten
	originalURL := "https://golang.org/"
	shortcode, err := store.ShortenURL(originalURL)
	require.NoError(t, err, "Unable to shorten URL: ", originalURL)

	// get full
	retrievedURL, err := store.GetFullURL(shortcode)
	require.NoError(t, err, "Unable to get full URL for shortcode: ", shortcode)

	// check that they match
	require.Equal(t, originalURL, retrievedURL)
}
