// Package dtos
package dtos

import "time"

// UpdateTaskInput representa la entrada para actualizar una tarea
type UpdateTaskInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   *bool     `json:"completed"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}
