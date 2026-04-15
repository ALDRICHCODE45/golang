// Package application
package application

import (
	dtos "hello_world/structs/application/dtos"
	domain "hello_world/structs/domain"
)

type FindAllUsersUseCase struct {
	userRepository domain.UserRepository
}

func NewFindAllUsersUseCase(userRepository domain.UserRepository) FindAllUsersUseCase {
	return FindAllUsersUseCase{userRepository: userRepository}
}

func (uc FindAllUsersUseCase) Execute() ([]dtos.UserResponseDTO, error) {
	users, err := uc.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	// Convertir []domain.User → []dtos.UserResponseDTO
	// Aquí es donde usas slices en Go
	result := make([]dtos.UserResponseDTO, 0, len(users)) // ← slice vacío con capacidad

	for _, user := range users {
		result = append(result, dtos.UserResponseDTO{
			ID:    user.ID,
			Name:  user.Name,
			Age:   user.Age.Value(),
			Email: user.Email.String(),
		})
	}

	return result, nil
}
