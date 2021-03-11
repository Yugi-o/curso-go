package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(tres(5))
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 23, 45, 68))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 46, 17, 25, 26, 98, 17))
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	}
	return 10, true
}

func tres(numero int) (z int) {
	z = numero * 2
	return
}

func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\nitem %d\n\t", item)
	}
	return total
}
