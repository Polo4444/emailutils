package otpstore

import (
	"fmt"
	"time"

	"github.com/Polo4444/emailutils/otp"
)

type InMemory struct {
	store map[string]otp.EmailCode
}

/* type Store interface {
	Register(code, purpose, entity string) (*EmailCode, error)
	List() ([]EmailCode, error)
	Get(code, purpose string) (*EmailCode, error)
	Remove(code, purpose) error
} */

// NewInMemory creates a new in-memory store
func NewInMemory() *InMemory {
	return &InMemory{
		store: map[string]otp.EmailCode{},
	}
}

// GenerateKeyString generates a new key string
func (i *InMemory) GenerateKeyString(code, purpose string) string {
	return fmt.Sprintf("%s|~|%s", code, purpose)
}

// Register registers a new code
func (i *InMemory) Register(code, purpose, entity string) (*otp.EmailCode, error) {
	em := &otp.EmailCode{
		CreatedAt: time.Now().UTC(),
		Purpose:   purpose,
		Code:      code,
		Entity:    entity,
	}
	i.store[i.GenerateKeyString(code, purpose)] = *em

	return em, nil
}

// List lists all the codes
func (i *InMemory) List() ([]otp.EmailCode, error) {

	codes := make([]otp.EmailCode, len(i.store))
	counter := 0
	for _, v := range i.store {
		codes[counter] = v
		counter++
	}
	return codes, nil
}

// Get gets a code. It returns an ErrCodeNotFound if the code is not found
func (i *InMemory) Get(code, purpose string) (*otp.EmailCode, error) {
	em, ok := i.store[i.GenerateKeyString(code, purpose)]
	if !ok {
		return nil, ErrCodeNotFound
	}
	return &em, nil
}

// Remove removes a code
func (i *InMemory) Remove(code, purpose string) error {
	delete(i.store, i.GenerateKeyString(code, purpose))
	return nil
}
