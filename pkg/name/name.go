package name

import (
	"errors"
	"fmt"
)

var ErrInvalidNameLength = errors.New("INVALID_NAME: name length must be more than 3")

type Name string

func IsValid(name string) (bool, error) {
	if len(name) < 3 {
		err := ErrInvalidNameLength
		return false, fmt.Errorf("input name : %s\n%w", name, err)
	}
	return true, nil
}
