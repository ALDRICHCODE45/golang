// Package infrastructure
package infrastructure

// TaskRepository define el contrato para el repositorio de tareas
// Esta es la interfaz que cualquier implementación (DynamoDB, in-memory, etc.) debe cumplir
import (
	"sync"
	"time"

	"github.com/aldrich/aws-practice/features/tasks/domain"
	"github.com/google/uuid"
)

// InMemoryTaskRepository es una implementación in-memory del repositorio
// Útil para desarrollo y testing
type InMemoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]domain.TaskEntity
}

// NewInMemoryTaskRepository crea una nueva instancia del repositorio in-memory
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]domain.TaskEntity),
	}
}

// Create crea una nueva tarea
func (r *InMemoryTaskRepository) Create(task *domain.TaskEntity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	r.tasks[task.ID] = *task
	return nil
}

// GetByID obtiene una tarea por su ID
func (r *InMemoryTaskRepository) GetByID(id string) (*domain.TaskEntity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, nil
	}
	return &task, nil
}

// GetAll obtiene todas las tareas
func (r *InMemoryTaskRepository) GetAll() ([]domain.TaskEntity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]domain.TaskEntity, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Update actualiza una tarea existente
func (r *InMemoryTaskRepository) Update(task *domain.TaskEntity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return nil
	}

	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = *task
	return nil
}

// Delete elimina una tarea por su ID
func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.tasks, id)
	return nil
}
