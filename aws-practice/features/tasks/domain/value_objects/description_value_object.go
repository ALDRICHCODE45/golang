// Package valueobjects
package valueobjects

import (
	"encoding/json"
	"errors"
	"strings"
)

type Description struct {
	value string
}

func NewDescription(value string) (Description, error) {
	cleanValue := strings.TrimSpace(value)

	if len(cleanValue) <= 5 {
		return Description{}, errors.New("la descripcion debe contener al menos 6 caracteres")
	}

	return Description{
		value: cleanValue,
	}, nil
}

func (description Description) Value() string {
	return description.value
}

func (description Description) MarshalJSON() ([]byte, error) {
	return json.Marshal(description.value)
}
