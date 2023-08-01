package main

import "fmt"

func main() {
	// GO es un lenguaje altamente tipado, aunque hay algunas excepciones.

	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Printf("%X", "pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("_________________________")

	// Declaracion de variables enteras
	base := 12 // crear y asignar; solo base = 12 es asignación para una variable ya creada
	var altura int = 14 // variable con tipo de valor
	var area int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)
	fmt.Println("_________________________")

	// Zero values
	var ancho int
	var largo float64
	var nombre string
	var casado bool

	fmt.Println("ancho:", ancho)
	fmt.Println("largo:", largo)
	fmt.Println("nombre:", nombre)
	fmt.Println("casado:", casado)
	fmt.Println("_________________________")


	var laBase float32
	fmt.Println("Ingrese la base:")
	fmt.Scan(&laBase)
	var laAltura float32
	fmt.Println("Ingrese la altura:")
	fmt.Scan(&laAltura)
	elArea := laBase * laAltura
	fmt.Println("area:", elArea)
}