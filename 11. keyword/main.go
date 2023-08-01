package main

import "fmt"

func main() {
	// Defer, se ejecuta al final. 
	// La buena practica es usar solo un Defer por función
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}