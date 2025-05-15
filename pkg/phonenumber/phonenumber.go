package phonenumber

import (
	"errors"
	"fmt"
)

// ErrInvalidphoneNumber is returned when the phone number is invalid.
var ErrInvalidphoneNumberLength = errors.New("INVALID_PHONE_NUMBER: phone number length must be 11")
var ErrInvalidphoneNumberDigits = errors.New("INVALID_PHONE_NUMBER: phone number must be digits")

func IsValid(phoneNumber string) (err error) {
	// TODO - we can use regex to validate phone number
	// phone number length validation
	if len(phoneNumber) < 11 || len(phoneNumber) > 11 {
		err = fmt.Errorf("input phnoe number : %s\n%w", phoneNumber, ErrInvalidphoneNumberLength)
		return
	}
	// phone number digits validation
	for _, c := range phoneNumber {
		if c < '0' || c > '9' {
			err = ErrInvalidphoneNumberDigits
			return
		}
	}
	return nil
}
