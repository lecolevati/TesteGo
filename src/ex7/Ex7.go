package main

import (
	"fmt"
	"math/rand"
)

func maiorLista(numeros ...int) int {
	var maior int
	cont := 0

	for _, numero := range numeros {
		if cont == 0 {
			maior = numero
		} else {
			if numero > maior {
				maior = numero
			}
		}
		cont++
	}

	return maior
}

func maiorVetor(vetorNumeros [100]int) (int, int) {
	var maior, posicao int
	cont := 0

	for _, numero := range vetorNumeros {
		if cont == 0 {
			maior = numero
			posicao = cont
		} else {
			if numero > maior {
				maior = numero
				posicao = cont
			}
		}
		cont++
	}

	return maior, posicao
}

func geraVetor() [100]int {
	var vetor [100]int

	for i := 0; i < 100; i++ {
		vetor[i] = int((rand.Float32() * 1000) + 1)
	}

	return vetor

}

func main() {
	var maiorDaLista, maiorDoVetor, posDoVetor int
	var vetor [100]int

	maiorDaLista = maiorLista(2, 18, 27, 35)
	fmt.Println("Maior da lista: ", maiorDaLista)

	vetor = geraVetor()
	fmt.Println("Vetor: ", vetor)

	maiorDoVetor, posDoVetor = maiorVetor(vetor)
	fmt.Println("vetor[", posDoVetor, "] = ", maiorDoVetor)
}
