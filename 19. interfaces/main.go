package main

import "fmt"

type figura2d interface {
	area() float32
}

type cuadrado struct {
	lado float32
}

type rectangulo struct {
	altura float32
	base float32
}

func (c *cuadrado) area() float32 {
	return c.lado * c.lado
}

func (r *rectangulo) area() float32 {
	return r.altura * r.base
}

func calcularArea (f figura2d) {
	fmt.Println( "Area: ", f.area() )
}

func main() {
	// USO DE INTERFACES
	myCuadrado := cuadrado{
		lado: 10,
	}

	myRectangulo := rectangulo{
		altura: 2,
		base:  4,
	}
	calcularArea(&myCuadrado)
	calcularArea(&myRectangulo)

	// LISTA DE INTERFACES
	// No se puede crear una lista de cualquier tipo, porque al
	// declararla, se tiene que mencionar el tipo. Para ello, se
	// hace la siguiente interface
	myInterface := []interface{}{
		"Hola",
		12,
		3.14,
	}
	fmt.Println(myInterface...)
}