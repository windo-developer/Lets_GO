package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("Not Found word")
var errCantUpdate = errors.New("Can't Update data")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
