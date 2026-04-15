// Package valueobjects
package valueobjects

import "errors"

type Age struct {
	value int
}

func NewAge(age int) (Age, error) {
	if age == 0 {
		return Age{}, errors.New("edad no valida")
	}
	if age <= 17 {
		return Age{}, errors.New("edad no valida")
	}

	return Age{value: age}, nil
}

func (age Age) Value() int {
	return age.value
}
