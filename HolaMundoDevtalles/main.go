package main

import (
	"fmt"
	"hello_world/structs/application"
	"hello_world/structs/application/dtos"
	"hello_world/structs/infrastructure"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. Crear repositorio en memoria
	userRepositoryInMemory := infrastructure.NewInMemoryUserRepository()

	// 2.Crear casos de uso inyectando el repositorio
	createUserUseCase := application.NewCreateUserUseCase(userRepositoryInMemory)
	findByIDUserUseCase := application.NewFindByIDUserUseCase(userRepositoryInMemory)
	findAllUsersUseCase := application.NewFindAllUsersUseCase(userRepositoryInMemory)
	deleteUserUseCase := application.NewDeleteUserUseCase(userRepositoryInMemory)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", r)

	// 3. crear usuarios

	user1, err := createUserUseCase.Execute(dtos.CreateUserDTO{
		Name:  "Aldrich",
		Age:   20,
		Email: "aldrich@aldrichcode.dev",
	})
	if err != nil {
		fmt.Println("error al crear usuario")
		return
	}

	fmt.Println("Primer Usuario creado: ", user1)

	user2, err := createUserUseCase.Execute(dtos.CreateUserDTO{
		Name:  "Byan",
		Age:   21,
		Email: "nocheblanca92@gmail.com",
	})
	if err != nil {
		fmt.Println("error al crear usuario2")
		return
	}
	fmt.Println("Primer Usuario2 creado: ", user2)

	userFounded, err := findByIDUserUseCase.Execute(user2.ID)
	if err != nil {
		fmt.Println("error al traer usuario")
		return
	}
	fmt.Println("Usuario encontrado", userFounded)

	allUsers, err := findAllUsersUseCase.Execute()
	if err != nil {
		fmt.Println("error al encontrar todos los usuarios")
		return
	}
	fmt.Println("Todos los usuarios", allUsers)

	err = deleteUserUseCase.Execute(user2.ID)
	if err != nil {
		fmt.Println("error al eliminar el usuario")
		return
	}

	fmt.Println("usuario eliminado")

	all, err := findAllUsersUseCase.Execute()
	if err != nil {
		fmt.Println("error al encontrar todos los usuarios")
		return
	}
	fmt.Println("Todos los usuarios", all)
}
