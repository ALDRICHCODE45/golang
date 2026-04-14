// Package valueobjects
package valueobjects

import (
	"errors"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, errors.New("email es requerido")
	}

	containAt := strings.Contains(email, "@")

	if !containAt {
		return Email{}, errors.New("email no es valido")
	}

	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}
