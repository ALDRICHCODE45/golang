// Package main
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"hello_world/structs/application"
	"hello_world/structs/application/dtos"
	"hello_world/structs/infrastructure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Event = string

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error codificando respuesta JSON", http.StatusInternalServerError)
	}
}

func main() {
	// 1. Crear repositorio en memoria
	userRepositoryInMemory := infrastructure.NewInMemoryUserRepository()

	// 2. Crear casos de uso inyectando el repositorio
	createUserUseCase := application.NewCreateUserUseCase(userRepositoryInMemory)
	findByIDUserUseCase := application.NewFindByIDUserUseCase(userRepositoryInMemory)
	findAllUsersUseCase := application.NewFindAllUsersUseCase(userRepositoryInMemory)
	updateUserUseCase := application.NewUpdateUserUseCase(userRepositoryInMemory)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := findAllUsersUseCase.Execute()
		if err != nil {
			writeJSON(w, http.StatusOK, []dtos.UserResponseDTO{})
			return
		}

		writeJSON(w, http.StatusOK, users)
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "id debe ser numérico"})
			return
		}

		user, err := findByIDUserUseCase.Execute(id)
		if err != nil {
			writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, user)
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var input dtos.CreateUserDTO

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "JSON inválido"})
			return
		}

		created, err := createUserUseCase.Execute(input)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, created)
	})

	r.Patch("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "id debe ser numérico"})
			return
		}

		var input dtos.UpdateUserDTO
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "Campos inválido"})
			return
		}

		if input.Name == nil && input.Age == nil && input.Email == nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: "debes enviar al menos un campo para actualizar"})
			return
		}

		updated, err := updateUserUseCase.Execute(id, input)
		if err != nil {
			if err.Error() == "usuario no encontrado" {
				writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
				return
			}

			writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, updated)
	})

	fmt.Println("Servidor corriendo en :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("error levantando servidor:", err)
	}
}
