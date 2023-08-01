package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping () {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram *= 2
}

func main() {
	a := 50
	b := &a // & es el operador de puntero para obtener la dirección en la memoria

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // * es el operador de dereferencia para obtener el valor en la memoria

	*b = 100
	fmt.Println(a)
	fmt.Println(b)

	myPc := pc{ram: 16, disk: 100, brand: "Lenovo"}
	myPc.ping()
	fmt.Println("Ram:", myPc.ram)
	myPc.duplicateRAM()
	fmt.Println("New RAM:", myPc.ram)
}