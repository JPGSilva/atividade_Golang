package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func remove(texto string) {

	vogais := []string{"a", "à", "á", "ã", "â",
		"e", "è", "é", "ê", "ẽ",
		"i", "í", "ì", "ĩ", "î",
		"o", "ò", "ó", "õ", "ô",
		"u", "ù", "ú", "ũ", "û"}

	var vogaisUpper []string

	for _, v := range vogais {
		v = strings.ToUpper(v)
		vogaisUpper = append(vogaisUpper, v)
	}

	for _, v := range vogais {
		texto = strings.ReplaceAll(texto, v, "")
		for _, v := range vogaisUpper {
			texto = strings.ReplaceAll(texto, v, "")
		}
	}
	fmt.Println("Frase 0  vogais: ", texto)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Conte-me um Poema: ")
	scanner.Scan()

	texto := scanner.Text()

	remove(texto)
}
