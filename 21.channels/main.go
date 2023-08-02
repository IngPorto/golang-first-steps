package main

import "fmt"

// Se puede indicar si el canal que se recibe como parámetro es un
// canal de entrada o de salida. Esto optimiza la memoria.
// De entrada: chan<- 
// De salida: <-chan
func say(text string, channel chan<- string) {
	// ingresar el dato al canal
	channel <- text

	// en caso que fuera solo de salida. Se extrae el dato del canal
	// y se asigna a la variable text:
	// text = <- channel
}
func main() {
	// creación del canal
	c := make(chan string, 2)

	fmt.Println("Hola")

	go say("Bye", c)

	// extraer el dato del canal. Sin exta linea no se imprime el "Bye"
	fmt.Println(<-c)

}