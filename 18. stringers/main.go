package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Para más detalles sobre los "vervos de formato", consulte la documentación
// https://pkg.go.dev/fmt#pkg-overview

// La estructura de datos " Struct " tiene un método llamado " String ", 
// que podemos sobrescribir para personalizar la salida a consola de los 
// datos del struct.
func (myPC pc) String() string {
	return fmt.Sprintf("Brand: %s, Ram: %d, Disk: %d", myPC.brand, myPC.ram, myPC.disk)
}

func main() {
	myPC := pc{ram: 16, disk: 100, brand: "Lenovo"}

	fmt.Println(myPC)
	// myPC.toString() // No se hace esto porque se está sobrescribiendo el método String
	fmt.Printf("%+v", myPC)
}