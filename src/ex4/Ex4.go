package main

import (
	"fmt"
)

func main() {
	var cod, qtd int
	var espec string
	var preco float32

	fmt.Print("Digite um codigo valido: ")
	fmt.Scan(&cod)
	fmt.Print("Digite a quantidade: ")
	fmt.Scan(&qtd)

	switch cod {
	case 100:
		espec = "Cachorro quente"
		preco = 1.2
	case 101:
		espec = "Bauru Simples"
		preco = 1.3
	case 102:
		espec = "Bauru com ovo"
		preco = 1.5
	case 103:
		espec = "Hamburger"
		preco = 1.2
	case 104:
		espec = "Cheeseburger"
		preco = 1.3
	case 105:
		espec = "Refrigerante"
		preco = 1.0
	default:
		preco = 0.0
	}

	if preco > 0 {
		valorTotal := float32(qtd) * preco
		fmt.Println(qtd, " - ", espec, " - Valor total: ", valorTotal)
	} else {
		fmt.Println("Codigo invalido")
	}
}
