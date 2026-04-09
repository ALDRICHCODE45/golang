package main

import "fmt"

func main() {
	evenNumbers := []int{3, 2, 3, 4, 5}
	fmt.Println("Hola Mundo", evenNumbers)

	for i, v := range evenNumbers {
		if v == 2 {
			fmt.Println("El numero: ", v, " Es dos", i)
		}
	}
}
