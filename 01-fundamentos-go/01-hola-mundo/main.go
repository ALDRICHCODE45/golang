package main

import "fmt"

func main() {
	name := "Aldrich"
	edad := 21
	learningGo := true
	weekHours := 5

	var city string = "Ciudad de Mexico"

	// bucle tipo for comun
	for i := 4; i < 5; i++ {
		fmt.Println("Tipo for normal")
	}

	// bucle tipo for comun
	suma := 100
	for i := 10; i > 5; i-- {
		suma -= 10
	}
	fmt.Println("Suma:", suma)

	evenNumbers := []int{}

	final := 0

	for final <= 10 {

		if final%2 == 0 {
			fmt.Println("Encontramos un numero par", final)
			evenNumbers = append(evenNumbers, final)
			final += 1
			continue
		}

		fmt.Println("No es numero par: ", final)

		final += 1
	}

	fmt.Println("Arreglo de numero pares final:", evenNumbers)

	// bucle tipo whileP
	n := 3
	for n < 5 {
		fmt.Println("N es menor a 5")
		n += 1
	}

	fmt.Println("Nombre:", name)
	fmt.Println("Edad:", edad)
	fmt.Println("Aprendiendo go?:", learningGo)
	fmt.Println("Horas a la semana?:", weekHours)
	fmt.Println("City?:", city)
}
