// Package application
package application

import "github.com/aldrich/aws-practice/features/tasks/domain"

// GetTaskByID obtiene una tarea por su ID
func (taskService *TaskService) GetTaskByID(id string) (*domain.TaskEntity, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}
	return taskService.repo.GetByID(id)
}
