package main

import "fmt"

func main() {
	// Operadores que permiten determinar si una comparación es
	// verdadera o falsa
	// COMPARACIÓN
	edadJuan := 20
	edadMaria := 21

	result := edadJuan == edadMaria // FALSE
	result = edadJuan != edadMaria	// TRUE
	result = edadJuan < edadMaria 	// TRUE
	result = edadJuan > edadMaria 	// FALSE
	result = edadJuan >= edadMaria 	// FALSE
	result = edadJuan <= edadMaria 	// TRUE

	fmt.Println(result)


	// LÓGICOS
	// utilizan puertas lógicas
	// AND
	result = edadJuan >= edadMaria && edadJuan <= edadMaria 
	// OR
	result = edadJuan >= edadMaria || edadJuan <= edadMaria
	// NOT
	result = !result

}