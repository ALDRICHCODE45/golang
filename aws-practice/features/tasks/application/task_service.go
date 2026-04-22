// Package application
package application

import (
	"errors"

	"github.com/aldrich/aws-practice/features/tasks/domain/repositories"
)

// ErrTaskNotFound se usa cuando no se encuentra la tarea
var ErrTaskNotFound = errors.New("task not found")

// ErrInvalidInput se usa cuando la entrada es inválida
var ErrInvalidInput = errors.New("invalid input")

// TaskService maneja la lógica de negocio relacionada con tareas
type TaskService struct {
	repo repositories.TaskRepository
}

// NewTaskService crea un nuevo servicio de tareas
func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}
