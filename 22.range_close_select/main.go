package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	channel := make(chan string, 2)
	channel <- "Mensaje 1"
	channel <- "Mensaje 2"

	fmt.Println(len(channel), cap(channel))

	// CLOSE: Al cerrar el canal no se permite ingresar más datos
	// incluso aunque haya capacidad. Es una buena práctica
	close(channel)

	// RANGE: iterar sobre cada uno de los valores del canal
	for msg := range channel {
		fmt.Println(msg)
	}

	// SELECT: Cuando empezamos a manejar múltiples canales y no tenemos 
	// certaza de qué canal va a responder primer

	email1 := make(chan string)
	email2 := make(chan string)
	// No sabemos cuál de los dos canales va a responder primero
	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)

	for i:=0; i<2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email recibido de email1", msg1)
		case msg2 := <-email2:
			fmt.Println("Email recibido de email2", msg2)
		}
	}
}