package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var vetor [100]int
	var contPares, contImpares int

	for i := 0; i < 100; i++ {
		vetor[i] = int((rand.Float32() * 1000) + 1)
	}
	fmt.Println(vetor)

	for i := 0; i < 100; i++ {

		if vetor[i]%2 == 0 {
			contPares++
		} else {
			contImpares++
		}

	}

	fmt.Println("Total pares: ", contPares)
	fmt.Println("Total impares: ", contImpares)

}
