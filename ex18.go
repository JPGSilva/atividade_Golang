package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cesar(texto string, param rune) string {
	for _, v := range strings.ToLower(texto) {
		if v <= 122 && v >= 97 {
			if v == 122 {
				v = 96
			}
			v = v + (param)
			fmt.Printf("%c", v)
		}
	}
	return texto
}

func main() {

	var param rune
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Digite uma Frase: ")
	scanner.Scan()
	texto = scanner.Text()

	fmt.Printf("Total Letras desvio: ")
	fmt.Scan(&param)

	for _, v := range strings.ToLower(texto) {
		fmt.Printf("%c", v)
	}

	fmt.Printf("\n\n")

	fmt.Println(cesar(texto, param))
}
