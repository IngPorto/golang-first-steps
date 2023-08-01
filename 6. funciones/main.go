package main

import "fmt"

func saludar(saludo string ){
	fmt.Println(saludo)
}

// a y b son del mismo tipo de argumentos
func tripleArgument (a , b int, c string) {
	fmt.Println(a, b, c)
}

func retrunSomeValue(a int) int{
	return a * 3
}

func doubleReturn(a int) (int, int){
	return a * 2, a * 3
}

func main() {
	saludar("Hola")
	saludar("Mundo")
	tripleArgument(1, 2, "Hola")
	retrunSomeValue(3)
	value1, value2 := doubleReturn(4)
	fmt.Println(value1, value2)
	// Tomar solo uno de los valores
	otroValue1, _ := doubleReturn(4)
	fmt.Println(otroValue1)
}