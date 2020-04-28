package main

import (
	"fmt"
	"math"
)

func imc(altura float64, peso float64) float64 {
	calculo := peso / math.Pow(altura, 2)
	return calculo
}

func main() {
	var altura, peso float64

	fmt.Print("Digite a altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Digite o peso: ")
	fmt.Scanln(&peso)

	imc := imc(altura, peso)
	fmt.Println("IMC: ", imc)

	imcFormatado := fmt.Sprintf("%.2f", imc)
	fmt.Println("IMC Formatado: ", imcFormatado)
}
