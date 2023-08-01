package main

// Las rutas del import deben ir desde la raiz del GOPATH
// si esto genera problemas, ejecutar el comando - go mod init -
// en la raiz de la carpeta del proyecto
import (
	"fmt"
	myAlias "helloGo/16.acceso/mypackage"
)

func main() {
	var myCar myAlias.CarPublic
	myCar.Brand = "Toyota"
	myCar.Year = 2020

	fmt.Println(myCar)

	// no se puede imprimir porque es privado
	// var otherCar myAlias.carPrivate 
	// fmt.Println(otherCar)

	myAlias.PrintMessage()
}