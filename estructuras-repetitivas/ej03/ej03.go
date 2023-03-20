package main

import "fmt"

func main() {
	valor := 0
	fmt.Print("Ingrese valor...  ")
	fmt.Scanln(&valor)
	fmt.Printf("¿%v es primo? %v", valor, CalcularNroPrimo(valor))
	//CalcularNroPrimo(valor)

}
func CalcularNroPrimo(n int) (resultado bool) {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
