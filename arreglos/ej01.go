package main

import "fmt"

func main() {
	//Positivos
	var array1 = []int{2, 45, 51, 2, 2}
	//Negativos
	var array2 = []int{-2, -45, -51, -2, -2}
	//Array vacío
	var array3 = []int{}
	fmt.Printf("Total array 1: %v \n", CalcularSumaArray(array1))
	fmt.Printf("Total array 2: %v \n", CalcularSumaArray(array2))
	fmt.Printf("Total array 3: %v \n", CalcularSumaArray(array3))
}
func CalcularSumaArray(n []int) (resultado int) {

	for i := 0; i < len(n); i++ {
		resultado += n[i]
	}
	return resultado

}
