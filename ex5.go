package main

import "fmt"

func medias(nota float64, nota2 float64) float64 {

	soma := nota2 / nota
	return soma
}

func main() {

	var cont int
	var notas int
	fmt.Printf("\n Quantas notas serão Adicionadas ? \n")
	fmt.Scan(&cont) // quantidade
	fmt.Printf("\n")
	fmt.Printf("Digite as Notas:")
	fmt.Scan(&notas) // notas

	n1 := float64(cont)
	n2 := float64(notas)

	for i := 1; i < cont; i++ {
		fmt.Printf("\n______________________________________________________________________________________________________________")

		fmt.Print("\nDigite a nota ", i, " : ")
		fmt.Println(notas)

		notas += notas
		fmt.Printf("\n______________________________________________________________________________________________________________")

	}

	fmt.Print("\n A media das notas é: ", medias(n1, n2))

}
