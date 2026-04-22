// Package application
package application

import (
	"github.com/aldrich/aws-practice/features/tasks/application/dtos"
	"github.com/aldrich/aws-practice/features/tasks/domain"
	valueobjects "github.com/aldrich/aws-practice/features/tasks/domain/value_objects"
)

// UpdateTask actualiza una tarea existente
func (taskService *TaskService) UpdateTask(id string, input dtos.UpdateTaskInput) (*domain.TaskEntity, error) {
	task, err := taskService.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, ErrTaskNotFound
	}

	// Actualizar solo los campos que vienen
	if input.Title != "" {
		titleVo, err := valueobjects.NewTitle(input.Title)
		if err != nil {
			return nil, err
		}
		task.Title = titleVo
	}

	if input.Description != "" {
		descriptionVO, err := valueobjects.NewDescription(input.Description)
		if err != nil {
			return nil, err
		}
		task.Description = descriptionVO
	}

	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if input.Priority != "" {
		task.Priority = input.Priority
	}

	if !input.DueDate.IsZero() {
		task.DueDate = input.DueDate
	}

	if err := taskService.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}
