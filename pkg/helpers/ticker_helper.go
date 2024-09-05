package helpers

import (
	"errors"
	"regexp"
)

// ValidateTicker validates the ticker symbol.
func ValidateTicker(ticker string) error {
	// Check if the ticker is empty
	if ticker == "" {
		return errors.New("ticker symbol must not be empty")
	}

	// Check if the ticker contains only alphabets and no special characters or spaces
	matched, err := regexp.MatchString("^[A-Za-z]+$", ticker)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("ticker symbol must only contain alphabets with no special characters or spaces")
	}

	// Check if the ticker length is within a reasonable range (1 to 5 characters)
	if len(ticker) < 1 || len(ticker) > 5 {
		return errors.New("ticker symbol must be between 1 and 5 characters long")
	}

	return nil
}
