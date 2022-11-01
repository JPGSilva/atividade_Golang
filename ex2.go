package main

import "fmt"

func main() {

	var ano int

	fmt.Printf(" Entre com o Ano para ser convertido em  Minutos e Segundos ! \n")
	fmt.Scanln((&ano))

	if ano <= 0 {
		fmt.Print("Ano Inválido por Favor tente novamente !  \n")

		fmt.Scanln((&ano))

	} else {

		min := ano * 525600

		seg := min * 60

		fmt.Println("\n", ano, "ano(s), equivale a : ", min, "minuto(s) e  ", seg, "segundos(s) ")
	}

}
