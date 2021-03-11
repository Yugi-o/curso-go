package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	var numero2, numero3, numero4 int
	numero5 := 15
	numero5 = 4
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	// No puedes cambiar el valor de una variable a otra de distinto tipo
	// var moneda float32 = 0
	// numero2 = moneda

	var moneda int64 = 0
	numero2 = int(moneda)
	// Convertir un numero a string
	// texto = fmt.Sprintf("%d", moneda)
	// strconv permite mas funciones que fmt para convertir datos a otros de distinto tipo
	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(texto)

	// Muestra el valoor por scoope
	fmt.Println(status)

	// Muestra el valor global
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
