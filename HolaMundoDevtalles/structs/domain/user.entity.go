// Package domain
package domain

import (
	"errors"

	vo "hello_world/structs/domain/value_objects"
)

type User struct {
	ID    int
	Name  string
	Age   vo.Age
	Email vo.Email
}

func NewUser(id int, email string, name string, Age int) (User, error) {
	if name == "" {
		return User{}, errors.New("name is invalid")
	}

	ageV0, err := vo.NewAge(Age)
	if err != nil {
		return User{}, err
	}

	emailV0, err := vo.NewEmail(email)
	if err != nil {
		return User{}, err
	}

	return User{ID: id, Name: name, Age: ageV0, Email: emailV0}, nil
}
