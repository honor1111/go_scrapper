package mydict

import (
	"errors"
)

// type Money int
// Money(1) == 1

// Dictionary for map
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errWordExists = errors.New("That word already exist")
	errCantUpdate = errors.New("That word does not exist")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, existance := d[word]
	if existance {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary only if word does not exist
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

// Update dictionary word as definition
func (d Dictionary) Update(word, defintion string) error {
	_, err := d.Search(word)

	switch err {

	case errNotFound:
		return errCantUpdate

	case nil:
		d[word] = defintion
	}

	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
