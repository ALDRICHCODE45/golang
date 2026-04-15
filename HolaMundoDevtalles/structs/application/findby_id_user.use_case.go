// Package application
package application

import (
	dtos "hello_world/structs/application/dtos"
	domain "hello_world/structs/domain"
)

type FindByIDUserUseCase struct {
	userRepository domain.UserRepository
}

func NewFindByIDUserUseCase(userRepository domain.UserRepository) FindByIDUserUseCase {
	return FindByIDUserUseCase{userRepository: userRepository}
}

func (uc FindByIDUserUseCase) Execute(id int) (dtos.UserResponseDTO, error) {
	user, err := uc.userRepository.FindByID(id)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	return dtos.UserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Age:   user.Age.Value(),
		Email: user.Email.String(),
	}, nil
}
