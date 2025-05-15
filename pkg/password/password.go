package password

import (
	"crypto/md5"
	"encoding/hex"
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
	// Define valid special characters
	validSpecialChars := []rune{'_', '-', '@'}

	// Convert password to runes for proper Unicode handling
	passwordRunes := []rune(password)

	// Check each character in the password
	for _, char := range passwordRunes {
		isValid := unicode.IsLetter(char) ||
			unicode.IsDigit(char)

		// Check if char is one of the allowed special characters
		for _, validChar := range validSpecialChars {
			if char == validChar {
				isValid = true
				break
			}
		}

		if !isValid {
			return false, INVALID_PASSWORD_CHARACTER
		}
	}

	return true, nil
}

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
