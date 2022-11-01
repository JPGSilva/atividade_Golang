package main

import "fmt"

func mult(n1, n2, n3 int) int {

	return n1 * n2 * n3
}

func main() {

	var n1, n2, n3 int

	fmt.Printf(" Entre com os  três Números para multiplicação: ")

	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)

	fmt.Println("O resultado da multiplicação dos 3 Números é: ", mult(n1, n2, n3))
}
