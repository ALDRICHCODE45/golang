// Package application
package application

// DeleteTask elimina una tarea
func (taskService *TaskService) DeleteTask(id string) error {
	if id == "" {
		return ErrInvalidInput
	}

	task, err := taskService.repo.GetByID(id)
	if err != nil {
		return err
	}

	if task == nil {
		return ErrTaskNotFound
	}

	return taskService.repo.Delete(id)
}
