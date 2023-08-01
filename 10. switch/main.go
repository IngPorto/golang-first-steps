package main

import "fmt"

func main() {
	switch modulo := 10 % 3; modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	// Switch sin ninguna condición, cuando se quiera anidad con 
	// multiples condiciones
	value := 200
	switch {
		case value > 100:
			fmt.Println("Es mayor que 100")
		case value < 100:
			fmt.Println("Es menor que 100")
		default:
			fmt.Println("Es igual a 100")
	}
}