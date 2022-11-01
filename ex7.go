package main

import "fmt"

func varredura(mediaNum, maior float64) (float64, float64) {

	var numero float64

	for i := 0; i < 10; i++ {

		fmt.Scan(&numero)
		if numero > maior {

			maior = numero
		}

		mediaNum += numero
	}

	mediaNum = mediaNum / 10

	return mediaNum, maior
}

func main() {

	var mediaNum, maior float64

	fmt.Println("Digite os 10 numeros abaixo:")

	mediaNum, maior = varredura(mediaNum, maior)

	fmt.Println("Média:", mediaNum)
	fmt.Println("O Maior Numero é:", maior)
}
