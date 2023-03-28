package main

import (
	"fmt"
	"sort"
)

func main() {
	array1 := []int{1, 2, 3, 5, 6}
	array2 := []int{1, 3, 5, 6, 7, 8}

	fmt.Printf("Interseccion: %v\n Union: %v", Interseccion(array1, array2), Union(array1, array2))
}

func Interseccion(a []int, b []int) (resultado []int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				resultado = append(resultado, a[i])
			}
		}
	}
	sort.Ints(resultado)
	return resultado
}

func Union(a []int, b []int) (resultado []int) {

	return append(a, b...)
}
