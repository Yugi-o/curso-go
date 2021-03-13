package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {

	}
	final := time.Now()
	canal <- final.Sub(inicio)
}
