package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func sayConWaitGroup(text string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // espera a que la función finalice y se ejecuta
	say(text)
}

func goroutineSimple() {
	say("Hola")
	// go envia la tarea a otro nucleo del procesador
	go say("Mundo")
	// paquete time para detener el proceso por unos segundos
	time.Sleep(1 * time.Second)
}

func goroutineConWaitGroup(){
	// Un WaitGroup agrupa goroutines para vigilar su finalización. Es mejor 
	// práctica que usar time.Sleep. Solo tiene 3 métodos: Add, Done y Wait
	wg := sync.WaitGroup{}

	say("Hola")
	// se crea un grupo de goroutines
	wg.Add(1)

	// Se envía una tarea junto con un WaitGroup para que el proceso envíe una
	// la señal de cuando se encuentra finalizada (Done)
	go sayConWaitGroup("Mundo", &wg)

	// Función anonima	
	go func(text string){
		text1, text2 := strings.Split(text, " ")[0], strings.Split(text, " ")[1]
		fmt.Println(text1)
		fmt.Println(text2)
	}("Vamos Golang")

	// Espera a que terminen todas las goroutines del WaitGroup
	wg.Wait()
}

func main() {
	// EJERCICIO 1
	goroutineSimple()

	// EJERCICIO 2
	goroutineConWaitGroup()
}