// Package dtos
package dtos

import "time"

// CreateTaskInput representa la entrada para crear una tarea
type CreateTaskInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}
