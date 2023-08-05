# First Steps on GO Lang

### 1. Hola Mundo en GO

### 2. Variables primeras líneas de código con Go

### 3. Estructuras de control de flujo y condicionales

### 4. Estructuras de datos básicas

### 5. Métodos e interfaces

### 6. Concurrencia y channels

### 7. Manejo de paquetes y Go Modules

## Recursos

* Gophers Slack: gophers.slack.com
* Go tour: tour.golang.org
* Golang Weekly: golangweekly.com
* Play With Go: play-with-go.dev
* Go by example: gobyexample.com
* Go Time (en Spotify)


## Cheat Sheet Go
https://devhints.io/go

* Hola mundo
```go
package main

import "fmt"

func main(){
    fmt.Println("Hola mundo")
}
```

* ¿Hacer una impresión en consola rápida?
```go
package main

func main(){
    print("Hola")
}
```

* Importar una librería sin usarla
```go
/*
    Hazlo solo y únicamente cuando la librería externa
    que estés usando lo pida explícitamente
*/
import ( 
    "fmt"
    _ "math"
)

func main(){
    fmt.Println("Hola mundo")
}
```

* Agregar un alias a un import (no suele usarse, pero es bueno saberlo)
```go
package main

import (
	"fmt"
	mth "math"
)

func main() {
	fmt.Println(mth.Pi)
}
```

* Diferentes formas de declarar variables
```go
v := 12
var v int = 12
var v int
```

* Zero values de primitivos
```go
var a int // 0
var b float64 // 0
var c string // ""
var d bool // false
```

* Incremental y decremental
```go
x++ // Suma 1 a x
x-- // Resta 1 a x
```

* Imprimir tipo de variables (hay otras formas, pero esta es la más fácil)
```go
a := 2
fmt.Printf("%T", a)
```

* Pedir información por consola
```go
a := fmt.Scanln(&variable)
```

* Función para tomar los errores (ahorra mucho código)
```go
func isError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Ejemplo de uso
func main() {
	_, err := strconv.Atoi("53a")
	isError(err)
}
```
* Arrays vs Slices
```go
// Array
var myList [2]int

// Slice
var myList2 []int
```

* Slice de interfaces (Úsalo con sabiduría)
```go
// Permite guardar diferentes tipos de datos en un mismo slice
myList := []interface{}{"Hola", 12, 4.90}

// Iterar sobre los distintos tipos de datos de ese slice
for _, v := range myList {
    switch v.(type) {
    case int:
        fmt.Println("Es int")
    case string:
        fmt.Println("Es string")
    case float64:
        fmt.Println("Es float64")
    }
}
```

* Asegurarnos si un key existe en el map
```go
m := make(map[string]int)

m["hola"] = 1

// Nota, usalmente se usa "ok" para recibir la segunda variable
value, ok := m["hello"]

/*
Si existe, ok será "true"
Si no existe, ok será "false"

En este caso, ok es "false" porque no existe.
*/
```

* Punteros
```go
a := 10 // Variable int
b := &a // "b" es el puntero de "a"
c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"
```

* Comandos de Go modules
```go
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod
```

## Librerías para desarrollo web con Go

### FrontEnd - Hugo
https://gohugo.io/

### BackEnd

* Echo: https://echo.labstack.com/
* Gin-Gonic: https://gin-gonic.com/
* Beego: https://beego.me/
* Revel: https://revel.github.io/
* Gorilla: https://www.gorillatoolkit.org/
* Buffalo: https://gobuffalo.io/en/

Extra
* Swaggo: 
https://github.com/swaggo/swag

## Data Science con Go

### Jupyter
https://jupyter.org/
https://github.com/cosmos72/gomacro
https://github.com/gopherdata/gophernotes

Jupyter es una de las principales herramientas que utilizamos los Data Scientists en el día a día ya que nos permite ejecutar código de manera fácil e iterativa pudiendo reciclar variabless en cualquier momento.

Al momento de instalarlo vía pip (manejador de paquetes para Python) viene instalado con el kernel de Python listo para ejecutar código Python.

A pesar, que Go es un lenguaje compilado, la comunidad ha creado un intérprete de Go llamado gomacro para ejecutar código sin compilar. Partiendo de ello, crearon también un kernel de Go para usarlo en Jupyter Notebook llamado gophernotes

## Manejo de DataFrames
https://github.com/tobgu/qframe
https://github.com/go-gota/gota
https://github.com/rocketlaunchr/dataframe-go

En este punto los más populares son: qframe, gota y dataframe-go.

Hasta la fecha, ninguno está en su versión estable, pero todos están haciendo un gran trabajo porque como lo notaste en el curso, en Go las variables vacías no son nulas sino que tienen un valor por defecto. Y esto en el mundo de los datos es todo un reto ya que tener datos nulos es el pan de cada día.

## Visualizaciones

* gonum/plot: Gonum no solo es el Numpy en Go sino que además tiene su propio código de visualización. En este caso gonum/plot te permite hacer visualizaciones estáticas.
https://github.com/gonum/plot
https://numpy.org/
* go-echarts: Para el caso de gráficos interactivos esta es una de las mejores opciones ya que además de los gráficos puedes crear tu propio dashboard con la misma librería.
https://github.com/go-echarts/go-echarts

## Machine Learning
* GoLearn: Tiene diferentes modelos, entre ellos lineales, regresiones y clasificación.
https://github.com/sjwhitworth/golearn
* Gorgonia: Es el más popular cuando requerimos implementar Deep Learning sino que además tiene la opción de implementar CUDA (hacer modelos de Deep Learning usando la tarjeta gráfica Nvidia)
https://github.com/gorgonia/gorgonia
https://platzi.com/cursos/ia/
https://developer.nvidia.com/CUDA-zone