package otpstore

import "fmt"

// Errors
var (
	// ErrCodeNotFound is returned when the code is not found
	ErrCodeNotFound = fmt.Errorf("code not found")
)
