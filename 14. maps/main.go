package main

import "fmt"

func crearMapConMake() map[int]string {
	m := make(map[int]string)
	m[1] = "Juan"
	m[2] = "Maria"
	return m
}

func crearMapSinMake() map[int]string {
	mapeador := map[int]string{1: "Juan", 2: "Maria"}

	// Al recorrer un map, no se devuelven los valores en el 
	// orden ingresado por ser un proceso CONCURRENTE,
	// para evitar esto, se debe usar -slice-
	for key, value := range mapeador {
		fmt.Println(key, value)
	}
	return mapeador
}

func crearMapConLlaveString() map[string]string {
	m := make(map[string]string)
	m["Juan"] = "Cali"
	m["Maria"] = "Medellin"
	return m
}

func crearMapVerificarLlave() map[string]int {
	m := make(map[string]int)
	m["Juan"] = 12
	m["Maria"] = 14

	// verifico que la llave exista, porque, aunque no exista, 
	// Go me devuelve el valor correspondiente a su ZeroValue.
	// Básicamente crea la llave con un valor por defecto
	if _, ok := m["JuanMuelas"]; ok {
		fmt.Println("JuanMuelas existe")
	} else {
		fmt.Println("JuanMuelas no existe")
	}
	return m
}

func main() {
	
}