package shortner

import (
	"errors"

	"github.com/juanefec/go-url-shortner/dbaccess"
)

// Shorten takes a url and stores it. It returns a shorter version
func Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("no empty strings as url")
	}
	id, e := dbaccess.StoreURL(url)
	if e != nil {
		return "", e
	}

	newURL := getNewURL(id)

	return newURL, nil
}

// GetOriginal takes an id to search for the original url
func GetOriginal(id string) (string, error) {
	original, e := dbaccess.GetURL(id)
	if e != nil {
		return "", e
	}
	return original, nil
}

func getNewURL(id string) string {
	return "http://localhost:4444/?i=" + id
}
