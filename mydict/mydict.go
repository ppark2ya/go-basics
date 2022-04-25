package mydict

import (
	"errors"
)

type Dictionary map[string]string

// var errNotFound = errors.New("not found")
// var errCantUpdate = errors.New("can't update non-existing word")
// var errWordExists = errors.New("that word already exists")
var (
	errNotFound = errors.New("not found")
	errCantUpdate = errors.New("can't update non-existing word")
	errWordExists = errors.New("that word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error)  {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
