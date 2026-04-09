package main

import (
	"errors"
	"fmt"
)

func sum(number1 int, number2 int) (int, error) {
	if number1 <= 0 || number2 <= 0 {
		return 0, errors.New("no puedes mandar un numero menor o igual a 0")
	}

	total := number1 + number2

	return total, nil
}

func countText(texto string) (int, error) {
	textrunes := []rune(texto)

	if len(textrunes) <= 0 {
		return 0, errors.New("no se puede mandar un string vacio")
	}

	count := 0
	for _, v := range textrunes {
		if v == ' ' {
			continue
		}
		count += 1
	}
	return count, nil
}

func main() {
	fmt.Println("Hola Mundo")

	name := "Aldrich es muy buen programador y va a sacar la certificacion"

	result, err := sum(99, 1)
	if err != nil {
		fmt.Println("Error detectado:", err)
		return
	}

	fmt.Println("No hay deteccion de Errorres", result)

	countResult, err := countText(name)
	if err != nil {
		fmt.Println("Error detectado en contar el texto:", err)
		return
	}

	fmt.Println("No hay deteccion de Errorres", countResult)
	fmt.Println("Este es el resultado sin funcion", len([]rune(name)))
}
