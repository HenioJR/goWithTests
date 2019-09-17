package maps

import "errors"

var (
	ErrUnknownWord  = errors.New("Sorry, unknown word")
	ErrRepeatedWord = errors.New("Sorry, repeated word")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] != "" {
		return d[key], nil
	}
	return "", ErrUnknownWord
}

//Map types are reference types, like pointers or slices
//because of this, Dictionary does not a pointer
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownWord:
		d[key] = value
	case nil:
		return ErrRepeatedWord
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) {
	d[key] = value
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
