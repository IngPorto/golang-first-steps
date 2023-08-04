package main

// Para poder intalar librerias en producción y código de terceros
// que no están en los paquetes core.

// RECORDAR, el archivo "go.mod" en la raíz del proyecto, no debe
// editarde de forma manual sino desde el comando "go mod"

// Ejerccio: Instalar un servidor HTTP
// https://echo.labstack.com/docs/quick-start

// Otros comandos:
// go mod verify // verificar si se ha instalado correctamente

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}