package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	edadJuan := 23
	edadMaria := 21

	if edadJuan == 23 {
		fmt.Println("Es 23")
	} else {
		fmt.Println("No es 23")
	}

	if edadJuan > edadMaria {
		fmt.Println("Juan es mayor que Maria")
	} else {
		fmt.Println("Juan es menor que Maria")
	}

	if edadJuan > 18 || edadMaria > 18 {
		fmt.Println("Juan y María son mayores de edad")
	} else {
		fmt.Println("Juan y María son menores de edad")
	}

	// Convertir un texto a un entero
	value, err := strconv.Atoi("123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)
}