package main

import "fmt"

func main() {

	array1 := []int{2, 45}
	array2 := []int{-2, -21}

	array3 := []int{4, -3}
	array4 := []int{5, 2}

	resultado1, resultado2 := CalcularSumaYProductoEscalar(array1, array2)
	resultado3, resultado4 := CalcularSumaYProductoEscalar(array3, array4)

	fmt.Printf("Total suma | producto escalar 1: %v | %v \n", resultado1, resultado2)
	fmt.Printf("Total suma | producto escalar 2:  %v | %v \n", resultado3, resultado4)
}
func CalcularSumaYProductoEscalar(n []int, n2 []int) (resultadoSuma int, resultadoProducto int) {
	//Suma
	for i := 0; i < len(n); i++ {
		resultadoSuma += n[i]
	}
	for j := 0; j < len(n2); j++ {

		resultadoSuma += n[j]
	}

	//Producto escalar

	resultadoProducto = n[1]*n2[1] + n[0]*n2[0]

	return resultadoSuma, resultadoProducto

}
