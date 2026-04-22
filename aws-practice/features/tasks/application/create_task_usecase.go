// Package application
package application

import (
	"github.com/aldrich/aws-practice/features/tasks/application/dtos"
	"github.com/aldrich/aws-practice/features/tasks/domain"
	valueobjects "github.com/aldrich/aws-practice/features/tasks/domain/value_objects"
)

// CreateTask crea una nueva tarea
func (taskService *TaskService) CreateTask(input dtos.CreateTaskInput) (*domain.TaskEntity, error) {
	if input.Title == "" {
		return nil, ErrInvalidInput
	}

	// Valores por defecto
	priority := input.Priority
	if priority == "" {
		priority = "medium"
	}

	titleVO, err := valueobjects.NewTitle(input.Title)
	if err != nil {
		return nil, err
	}

	descriptionVO, err := valueobjects.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

	task := &domain.TaskEntity{
		Title:       titleVO,
		Description: descriptionVO,
		Priority:    priority,
		DueDate:     input.DueDate,
		Completed:   false,
	}

	if err := taskService.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}
