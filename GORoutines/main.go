package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLentooo("Yugi")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLentooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
