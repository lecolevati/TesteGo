package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 float32

	fmt.Print("Digite um numero: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite um numero: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite um numero: ")
	fmt.Scanln(&n3)

	media := (n1 + n2 + n3) / 3

	fmt.Println("Média das notas do aluno: ", media)
	fmt.Print("Média das notas do aluno com 2 digitos: ")
	mediaDoisDig := fmt.Sprintf("%.2f", media)
	fmt.Print(mediaDoisDig)

}
