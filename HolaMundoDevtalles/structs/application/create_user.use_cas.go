// Package application
package application

import (
	dtos "hello_world/structs/application/dtos"
	domain "hello_world/structs/domain"
)

type CreateUserUseCase struct {
	userRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) CreateUserUseCase {
	return CreateUserUseCase{userRepository: userRepository}
}

func (uc CreateUserUseCase) Execute(createUserDto dtos.CreateUserDTO) (dtos.UserResponseDTO, error) {
	user, err := domain.NewUser(0, createUserDto.Email, createUserDto.Name, createUserDto.Age)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	savedUser, err := uc.userRepository.Save(user)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	return dtos.UserResponseDTO{
		ID:    savedUser.ID,
		Name:  savedUser.Name,
		Age:   savedUser.Age.Value(),
		Email: savedUser.Email.String(),
	}, nil
}
