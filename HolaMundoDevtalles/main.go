package main

import (
	"fmt"
	"hello_world/utils"
)

func main() {
	fmt.Println("Hola Mundo")

	name := "Aldrich es muy buen programador y va a sacar la certificacion"
	edad := 18

	result, err := utils.Sum(99, 1)
	if err != nil {
		fmt.Println("Error detectado:", err)
		return
	}

	isAdult, err := utils.CalculateAdult(edad)

	if err != nil {
		fmt.Println("Error detectado calculando la edad del usuario:", err)
		return
	}

	fmt.Println("El usuario es adulto: ", utils.Booltolabel(isAdult, "SI", "NO"))

	fmt.Println("No hay deteccion de Errorres", result)

	countResult, err := utils.CountText(name)
	if err != nil {
		fmt.Println("Error detectado en contar el texto:", err)
		return
	}

	fmt.Println("No hay deteccion de Errorres", countResult)
	fmt.Println("Este es el resultado sin funcion", len([]rune(name)))
}
