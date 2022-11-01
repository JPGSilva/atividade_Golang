package main

import "fmt"

func somar(nota1 int, nota2 int) int {

	soma := nota1 + nota2
	return soma
}

func main() {

	var num1 int
	var num2 int

	fmt.Printf("\nAgora, digite o valor da primeira nota para fazer a media \n")
	fmt.Scanln(&num1)

	fmt.Printf("\n\n")
	fmt.Printf("\nAgora, digite o valor da segunda nota para fazer a media \n")
	fmt.Scanln(&num2)

	fmt.Printf("\n\n")
	fmt.Println("A soma dos valores é: ", somar(num1, num2))

}
