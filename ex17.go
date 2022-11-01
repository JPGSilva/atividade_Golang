package main

import (
	"fmt"
	"sort"
)

func main() {

	listaNomes := []string{"Kayblack",
		"Hariel",
		"Ryan",
		"Kyan",
		"Dfideliz",
		"Poze",
		"Igor"}

	fmt.Println("Nomes Escolhidos:", listaNomes)

	sort.Strings(listaNomes)
	fmt.Println("Ordem Alfabética:", listaNomes)
}
