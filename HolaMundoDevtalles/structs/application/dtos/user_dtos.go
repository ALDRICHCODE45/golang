// Package dtos
package dtos

type CreateUserDTO struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UserResponseDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UpdateUserDTO struct {
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Email *string `json:"email"`
}
