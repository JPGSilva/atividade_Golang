package main

import "fmt"

func main() {

	var x float64

	x = 100
	fmt.Print("\n Digite o total a ser Convertido : R$  ")
	fmt.Scanln(&x)

	conversor := x / 5.17
	fmt.Printf("  %0.2f", conversor)

}
