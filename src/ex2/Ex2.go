package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Print("Digite um numero: ")
	fmt.Scan(&n1)
	fmt.Print("Digite um numero: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Print("Maior: ")
		fmt.Println(n1)
	} else if n2 > n1 {
		fmt.Print("Maior: ")
		fmt.Println(n2)
	} else {
		fmt.Print("Numeros iguais")
	}

}
