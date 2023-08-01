package main

import "fmt"

func main() {
	// Array
	// La diferencia principal entre los arrays es que estos 
	// tienen una longitud fija e invariable y deben declarase 
	// especifiandola.
	// Son inmutables, su length y capacity serán siempre las mismas.
	var array [3]int
	array[0] = 1
	array[1] = 2

	// len() devuelve el tamaño del array
	// cap() devuelve la capacidad maxima del array

	fmt.Println(array, len(array), cap(array))

	// Slice
	// mientras que los Slices tienen una longitud variable y no 
	// hay que especificarla en la declaración. 
	// Debido a que no son inmutables y su tamaño es dinámico
	var slice []int
	slice = []int{0, 1, 2, 3, 4}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	slice = append(slice, 5)
	fmt.Println(slice)
	
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4]) // slice[low : high]
	fmt.Println(slice[4:])


	// Si queremos crear un slice deberiamos usar la funcion make:
	x := make([]float64, 5)
	// esto crea un Slice asociado a un array subjacente de longitud 5.
	// Los Slices siempre están asociados a un array y aunque nunca 
	// pueden ser mas largos que el aray, pueden ser mas cortos.
	// La función make también permite un tercer parámetro, que representa 
	//la capacidad del array, por lo que
	x = make([]float64, 5, 10)
	// representa un Slice de longitud 5 y capacidad de 10
	fmt.Println(x)

	// Agregar una lista de elementos
	newSlice := []int{1, 2, 3, 4, 5}
	slice = append(slice, newSlice...)
	fmt.Println("Slice con otra lista ", slice)
}