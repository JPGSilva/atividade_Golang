package main

import "fmt"

func main() {

	fmt.Printf("\nDigite a nota do cliente: ")
	var nota int

	fmt.Scanln(&nota)

	if nota >= 0 && nota < 7 {

		fmt.Printf("\nA avaliação desse cliente é: DETRATOR")

	}

	if nota > 6 && nota < 9 {

		fmt.Printf("\nA avaliação desse cliente é: NEUTRO")

	}

	if nota > 8 && nota <= 10 {

		fmt.Printf("\nA avaliação desse cliente é: PROMOTOR ")

	}

	if nota < 0 || nota > 10 {

		fmt.Printf("\nAvaliação Fora da Métrica\n")

	}
}
