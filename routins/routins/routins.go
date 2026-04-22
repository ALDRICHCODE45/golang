// Package routines
package routines

import "fmt"

func ingresarString(canal chan<- string) {
	canal <- "Hola Mundo"
	fmt.Println("Mensaje ingresado")
}

func printChannelMessage(canal <-chan string) {
	msg := <-canal

	fmt.Println(msg)
}

func Canales() {
	channel := make(chan string)

	go ingresarString(channel)
	printChannelMessage(channel)

	channel2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- i
		}

		close(channel2)
	}()

	for num := range channel2 {
		fmt.Println("Numero ingresado", num)
	}
}
