package main

// 1. Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes),
//  muestre por pantalla el polinomio completo, por ejemplo si recibe los coeficientes 3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2

import (
	"fmt"
	"strconv"
)

func MostrarCoeficientes(n ...float32) {
	x := 0
	total := ""
	for _, valor := range n {

		if valor != 0 {

			if x >= 2 {

				if x == len(n)-1 {
					total += (strconv.FormatFloat(float64(valor), 'f', 1.0, 32) + "X**" + strconv.Itoa(x) + " ")
				} else {
					total += (strconv.FormatFloat(float64(valor), 'f', 1.0, 32) + "X**" + strconv.Itoa(x) + " + ")
				}

			}

			if x == 1 {
				total += (strconv.FormatFloat(float64(valor), 'f', 1.0, 32) + "X + ")
			}

			if x == 0 {
				total += (strconv.FormatFloat(float64(valor), 'f', 1.0, 32) + " + ")
			}
		}
		x++
	}
	fmt.Print(total + "\n")

}
func main() {
	//Con coeficiente negativo
	MostrarCoeficientes(-2.0, 3.0, 5.0)
	//Con coeficiente Cero
	MostrarCoeficientes(0.0, 3.0, 0.0, 4.2, 2.3, 4.5)
	//Con coeficiente positivos
	MostrarCoeficientes(14.0, 13.0, 8.0)
	//Con muchos coeficientes
	MostrarCoeficientes(15.0, 7.0, 5.0, 0.0, 7.3, -4.0, -2.0, 42.0)
}
