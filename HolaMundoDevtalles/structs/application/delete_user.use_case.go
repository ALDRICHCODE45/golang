// Package application
package application

import (
	"hello_world/structs/domain"
)

type DeleteUserUseCase struct {
	userRepository domain.UserRepository
}

func NewDeleteUserUseCase(userRepository domain.UserRepository) DeleteUserUseCase {
	return DeleteUserUseCase{userRepository: userRepository}
}

func (uc DeleteUserUseCase) Execute(id int) error {
	_, err := uc.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	return uc.userRepository.Delete(id)
}
