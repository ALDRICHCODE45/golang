// Package infrastructure
package infrastructure

import (
	"errors"

	"hello_world/structs/domain"
)

type InMemoryUserRepository struct {
	users []domain.User
	nextD int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []domain.User{},
		nextD: 1,
	}
}

func (repo *InMemoryUserRepository) Save(user domain.User) (*domain.User, error) {
	user.ID = repo.nextD

	repo.nextD++

	repo.users = append(repo.users, user)

	return &user, nil
}

func (repo *InMemoryUserRepository) FindByID(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("id ingresado no valido")
	}

	for i := range repo.users {
		if repo.users[i].ID == id {
			return &repo.users[i], nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}

func (repo *InMemoryUserRepository) FindAll() ([]domain.User, error) {
	sliceLength := len(repo.users)

	if sliceLength == 0 {
		return nil, errors.New("no hay usuarios")
	}

	return repo.users, nil
}

func (repo *InMemoryUserRepository) Delete(id int) error {
	sliceLength := len(repo.users)

	if sliceLength == 0 {
		return errors.New("no hay usuarios")
	}

	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}

	return errors.New("usuario con ID no encontrado")
}

func (repo *InMemoryUserRepository) Update(user domain.User) (*domain.User, error) {
	if user.ID <= 0 {
		return nil, errors.New("id ingresado no valido")
	}

	for i := range repo.users {
		if repo.users[i].ID == user.ID {
			repo.users[i] = user
			return &repo.users[i], nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}
