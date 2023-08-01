package main

import "fmt"

// type es el Keyword para definir un tipo de dato.
// Un struct es simplemente una colección de campos.
type car struct {
	brand string
	year int
	parts map[string]int
}

func main (){
	myCar := car{
		brand: "Toyota",
		year: 2020,
		parts: map[string]int{
			"engine": 1,
			"transmission": 1,
			"doors": 5,
		},
	}

	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Ford"

	fmt.Println(otherCar)
}