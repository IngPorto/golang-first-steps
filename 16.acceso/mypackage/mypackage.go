package mypackage

import "fmt"

// CarPublic es una estructura publica
type CarPublic struct { // struct en mayúsculas son públicos
	Brand string					// atributos en mayúsculas son públicos
	Year int
}

type carPrivate struct {
	brand string
	year int
}

func PrintMessage () {
	fmt.Println("Hola, este es un mensaje")
}