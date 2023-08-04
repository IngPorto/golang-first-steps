/**
En esta clase se clona un módulo de Go directamente desde 
github. Luego se modifica go.mod y se instala el módulo
descargado. Esto se hace para poder editar los archivos
del modulo, porque instalandolo con "go get" o "go install"
los archivos quedan solo de lectura.

---
INSTALAR PAQUETE EDITABLE
Clono mi repo
> git clone https://github.com/labstack/echo

Reemplazo el repo previamente instalado por el clonado (el cual puedo manipular a mi antojo)
> go mod edit -replace github.com/labstack/echo/v4=./echo

Esto modifica el archivo go.mod creando la siguiente línea:
edit -replace github.com/labstack/echo/v4 =>./echo

---
REVERTIR
Cuando quiero revertir lo hecho, se debe ejecutar:
> go mod edit -dropreplace github.com/labstack/echo/v4

---
EMPAQUETAR
Empaquetar código de terceros en tu proyecto. Porque nativamente
los módulos, librerías y paquetes de terceros se guardan
en el GOPATH/pkg y no en mi proyecto. Así que cuando vaya a 
desplegar mi proyecto en producción, no se llevaría las dependencias

> go mod vendor

---
LIMPIAR PAQUETES NO USADOS
go.sum hacer una compilación de todos los paquetes usados

> go mod tidy

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}