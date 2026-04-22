// Package valueobjects
package valueobjects

import (
	"encoding/json"
	"errors"
	"strings"
)

type Title struct {
	value string
}

func NewTitle(value string) (Title, error) {
	cleanValue := strings.TrimSpace(value)

	if len(cleanValue) <= 3 {
		return Title{}, errors.New("el titulo debe ser mayor a 3 characteres")
	}

	return Title{
		value: cleanValue,
	}, nil
}

func (title Title) Value() string {
	return title.value
}

func (title Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(title.value)
}
