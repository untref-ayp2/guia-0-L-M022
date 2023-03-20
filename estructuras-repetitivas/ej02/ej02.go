package main

import "fmt"

func main() {
	fmt.Printf("Producto de 5: %v \n", CalcularProducto(5, 2))
	fmt.Printf("Producto de 12: %v \n", CalcularProducto(12, 3))
	fmt.Printf("Producto de 2: %v \n", CalcularProducto(2, 9))
	fmt.Printf("Producto de 0: %v \n", CalcularProducto(0, 1))
}
func CalcularProducto(a int, b int) (resultado int) {
	if b < 1 {
		return b
	}
	return a + CalcularProducto(a, b-1)

}
