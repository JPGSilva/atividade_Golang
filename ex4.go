package main

import "fmt"

func somaNotas(nota1, nota2 int8) int8 {
	return nota1 + nota2
}

func main() {

	var nota1, nota2 int8

	fmt.Printf("Entre com as Duas notas : ")
	fmt.Scanf("%d", &nota1)
	fmt.Scanf("%d", &nota2)

	fmt.Println("\n\n Resultado é :", somaNotas(nota1, nota2))

}
