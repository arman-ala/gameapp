package password

import (
	"fmt"
	"unicode"
)

const (
	MinLength = 8
	MaxLength = 32
)

var (
	INVALID_PASSWORD_LENGTH    = fmt.Errorf("invalid password length (8 to 32)")
	INVALID_PASSWORD_CHARACTER = fmt.Errorf("invalid password character (a-z 0-9 @ - _)")
)

func IsValid(password string) (valid bool, err error) {
	if len(password) < MinLength || len(password) > MaxLength {
		return false, INVALID_PASSWORD_LENGTH
	}

	// check password characters
	for i := 0; i < len(password); {
		if unicode.IsDigit(rune(password[i])) == true || unicode.IsLetter(rune(password[i])) == true || rune(password[i]) == '_' || rune(password[i]) == '-' || rune(password[i]) == '@' {
			i++
		} else {
			return false, INVALID_PASSWORD_CHARACTER
		}
	}

	return true, nil
}
