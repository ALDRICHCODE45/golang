// Package application
package application

import "github.com/aldrich/aws-practice/features/tasks/domain"

// GetAllTasks obtiene todas las tareas
func (taskService *TaskService) GetAllTasks() ([]domain.TaskEntity, error) {
	return taskService.repo.GetAll()
}
