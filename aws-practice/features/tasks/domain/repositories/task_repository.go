// Package repositories
package repositories

import "github.com/aldrich/aws-practice/features/tasks/domain"

// TaskRepository define el contrato para el repositorio de tareas
// Cualquier implementación (DynamoDB, in-memory, PostgreSQL, etc.) debe cumplir esta interfaz
type TaskRepository interface {
	Create(task *domain.TaskEntity) error
	GetByID(id string) (*domain.TaskEntity, error)
	GetAll() ([]domain.TaskEntity, error)
	Update(task *domain.TaskEntity) error
	Delete(id string) error
}
