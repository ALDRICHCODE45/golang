// Package domain
package domain

type UserRepository interface {
	FindByID(id int) (*User, error)
	FindAll() ([]User, error)
	Save(user User) (*User, error)
	Delete(id int) error
}
