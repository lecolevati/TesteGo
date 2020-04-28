package main

import (
	"fmt"
)

func main() {
	var num, maior int

	for i := 0; i < 3; i++ {
		fmt.Print("Digite um numero: ")
		fmt.Scanln(&num)

		if i == 0 {
			maior = num
		} else {
			if num > maior {
				maior = num
			}
		}
	}

	fmt.Println("Maior: ", maior)

	j := 0
	for j < 3 {
		fmt.Print("Digite um numero: ")
		fmt.Scanln(&num)

		if j == 0 {
			maior = num
		} else {
			if num > maior {
				maior = num
			}
		}
		j++
	}

	fmt.Println("Maior: ", maior)
}
