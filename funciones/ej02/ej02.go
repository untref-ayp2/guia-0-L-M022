package main

import (
	"fmt"
)

func main() {
	opciones := [4]string{"Correcta", "Incorrecta", "Correcta", "Incorrecta"}
	valor := 0

	fmt.Print("\n	Opción 1 \n 	Opción 2 \n 	Opción 3 \n 	Opción 4 \n")

	fmt.Print("\n> Ingrese opción: ")
	fmt.Scanln(&valor)

	if valor > 4 || valor < 1 {
		fmt.Print("\n Error: opción fuera de los limites \n")
		return
	} else {
		fmt.Printf("Opción %v: "+opciones[valor-1], valor)
	}
}
