package main

import (
	"fmt"
)

func ObtenerMenorYMayor(n ...int) {
	menor := 0
	mayor := 0

	for _, valor := range n {
		if menor > valor {
			menor = valor
		}
		if mayor < valor {
			mayor = valor
		}
	}

	fmt.Printf("Numero menor: %v \nNumero mayor: %v", menor, mayor)
}

func main() {
	ObtenerMenorYMayor(2, 4, 5, 6, 1, 6, 2, 8, 984, 32, 1, -421)
}
