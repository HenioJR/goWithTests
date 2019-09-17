package maps

import "errors"

var ErrUnknownWord = errors.New("Sorry, unknown word")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] != "" {
		return d[key], nil
	}
	return "", ErrUnknownWord
}
