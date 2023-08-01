package main

import "fmt"

func main() {
	x := 5
	y := 10
	
	// Suma
	fmt.Println("Suma: ", x + y)
	
	// Resta
	fmt.Println("Resta: ", x - y)
	
	// Multiplicacion
	fmt.Println("Multiplicacion: ", x * y)

	// Division
	fmt.Println("Division: ", x / y)

	// Módulo / Residuo
	fmt.Println("Modulo: ", x % y)

	// Incremento
	x++
	fmt.Println("Incremento: ", x)

	// Decremento
	x--
	fmt.Println("Decremento: ", x)

	fmt.Println("Hello World!")
}