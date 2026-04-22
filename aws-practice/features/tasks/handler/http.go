// Package handler
package handler

import (
	"errors"
	"net/http"

	"github.com/aldrich/aws-practice/features/tasks/application"
	"github.com/aldrich/aws-practice/features/tasks/application/dtos"
	"github.com/gin-gonic/gin"
)

// TaskHandler maneja los endpoints HTTP para tareas
type TaskHandler struct {
	service *application.TaskService
}

// NewTaskHandler crea un nuevo handler de tareas
func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// RegisterRoutes registra las rutas de tareas en el router
func (h *TaskHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/tasks", h.GetAllTasks)
	router.GET("/tasks/:id", h.GetTaskByID)
	router.POST("/tasks", h.CreateTask)
	router.PUT("/tasks/:id", h.UpdateTask)
	router.DELETE("/tasks/:id", h.DeleteTask)
}

// GetAllTasks obtiene todas las tareas
// @Summary Listar todas las tareas
// @Tags tasks
// @Produce json
// @Success 200 {array} domain.Task
// @Router /tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

// GetTaskByID obtiene una tarea por su ID
// @Summary Obtener tarea por ID
// @Tags tasks
// @Produce json
// @Param id path string true "ID de la tarea"
// @Success 200 {object} domain.Task
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// CreateTask crea una nueva tarea
// @Summary Crear tarea
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body application.CreateTaskInput true "Tarea a crear"
// @Success 201 {object} domain.Task
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input dtos.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := h.service.CreateTask(input)
	if err != nil {
		if err == application.ErrInvalidInput {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": task,
	})
}

// UpdateTask actualiza una tarea existente
// @Summary Actualizar tarea
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID de la tarea"
// @Param task body application.UpdateTaskInput true "Datos a actualizar"
// @Success 200 {object} domain.Task
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var input dtos.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := h.service.UpdateTask(id, input)
	if err != nil {
		if errors.Is(err, application.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, application.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// DeleteTask elimina una tarea
// @Summary Eliminar tarea
// @Tags tasks
// @Param id path string true "ID de la tarea"
// @Success 204
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteTask(id)
	if err != nil {
		if errors.Is(err, application.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
