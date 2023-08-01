package main

import (
	"fmt"
)

func main() {
	// Declaración de variables
	hellowMessage := "Hello"
	worldMessage := "World"

	fmt.Println(hellowMessage, worldMessage)
	fmt.Println(hellowMessage, worldMessage)

	nombre := "Juan"
	edad := 20
	// %s es string, %d es int
	fmt.Printf("%s tiene %d años\n", nombre, edad)
	// si no sabemos qué tipo de dato viene en la variable, usamos %v
	fmt.Printf("%v tiene %v años\n", nombre, edad)

	// para guardar un string desde la consola en una variable Sprintf
	message := fmt.Sprintf("%s tiene %d años", nombre, edad)
	fmt.Println(message)

	// Para conocer el tipo de variable usamos %T
	fmt.Printf("nombre %T\n", nombre)
	fmt.Printf("edad %T\n", edad)

	//// TODAS https://golang.org/pkg/fmt/
}