// Package domain
package domain

import (
	"time"

	valueobjects "github.com/aldrich/aws-practice/features/tasks/domain/value_objects"
)

type TaskEntity struct {
	ID          string                   `json:"id" dynamodbav:"id"`
	Title       valueobjects.Title       `json:"title" dynamodbav:"title"`
	Description valueobjects.Description `json:"description" dynamodbav:"description"`
	Completed   bool                     `json:"completed" dynamodbav:"completed"`
	Priority    string                   `json:"priority" dynamodbav:"priority"` // low, medium, high
	DueDate     time.Time                `json:"dueDate" dynamodbav:"dueDate"`
	CreatedAt   time.Time                `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt" dynamodbav:"updatedAt"`
	Attachments []string                 `json:"attachments,omitempty" dynamodbav:"attachments,omitempty"`
}
