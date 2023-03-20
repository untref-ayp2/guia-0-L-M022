package main

import "fmt"

func main() {
	fmt.Printf("Factorial de 5: %v \n", CalcularFactorial(5))
	fmt.Printf("Factorial de 12: %v \n", CalcularFactorial(12))
	fmt.Printf("Factorial de 2: %v \n", CalcularFactorial(2))
	fmt.Printf("Factorial de 0: %v \n", CalcularFactorial(0))
}
func CalcularFactorial(n int) (resultado int) {
	if n <= 1 {
		return n
	}
	return n * CalcularFactorial(n-1)

}
