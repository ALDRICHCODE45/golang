// Package application
package application

import (
	"errors"

	"hello_world/structs/application/dtos"
	"hello_world/structs/domain"
)

type UpdateUserUseCase struct {
	userRepository domain.UserRepository
}

func NewUpdateUserUseCase(userRepository domain.UserRepository) UpdateUserUseCase {
	return UpdateUserUseCase{userRepository: userRepository}
}

func (uc UpdateUserUseCase) Execute(id int, updateUserDto dtos.UpdateUserDTO) (dtos.UserResponseDTO, error) {
	user, err := uc.userRepository.FindByID(id)
	if err != nil {
		return dtos.UserResponseDTO{}, errors.New("usuario no encontrado")
	}

	name := user.Name
	if updateUserDto.Name != nil {
		name = *updateUserDto.Name
	}

	age := user.Age.Value()
	if updateUserDto.Age != nil {
		age = *updateUserDto.Age
	}

	email := user.Email.String()
	if updateUserDto.Email != nil {
		email = *updateUserDto.Email
	}

	userToUpdate, err := domain.NewUser(user.ID, email, name, age)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	updatedUser, err := uc.userRepository.Update(userToUpdate)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	return dtos.UserResponseDTO{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Age:   updatedUser.Age.Value(),
		Email: updatedUser.Email.String(),
	}, nil
}
