package main

import (
	"figuras/figuras"
	"fmt"
	"strconv"
	"strings"
)

func mostrar(f figuras.Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
	fmt.Println("\n")
}

func crearRectangulo() figuras.Rectangulo {
	valor := ""
	fmt.Print("Creando Rectangulo... \n Inserte coordenadas X,Y para el Primer punto... \n")
	fmt.Scanln(&valor)

	valores := strings.Split(valor, ",")
	x1, _ := strconv.ParseFloat(valores[0], 64)
	y1, _ := strconv.ParseFloat(valores[1], 64)
	p1 := figuras.Punto{X: x1, Y: y1}

	fmt.Print(" Inserte coordenadas X,Y para el Segundo punto... \n")
	fmt.Scanln(&valor)

	valores = strings.Split(valor, ",")
	x1, _ = strconv.ParseFloat(valores[0], 64)
	y1, _ = strconv.ParseFloat(valores[1], 64)
	p2 := figuras.Punto{X: x1, Y: y1}

	r := figuras.Rectangulo{P1: p1, P2: p2}
	return r
}

func crearCuadrado() figuras.Cuadrado {
	valor := ""
	fmt.Print("Creando Cuadrado... \n Inserte coordenadas X,Y para el punto... \n")
	fmt.Scanln(&valor)

	valores := strings.Split(valor, ",")
	x1, _ := strconv.ParseFloat(valores[0], 64)
	y1, _ := strconv.ParseFloat(valores[1], 64)
	p1 := figuras.Punto{X: x1, Y: y1}

	fmt.Print(" Inserte longitud de Lado... \n")
	fmt.Scanln(&valor)

	lado, _ := strconv.ParseFloat(valor, 64)

	r := figuras.Cuadrado{Pto: p1, Lado: lado}
	return r
}

func crearCirculo() figuras.Circulo {
	valor := ""
	fmt.Print("Creando Circulo... \n Inserte coordenadas X,Y para el centro... \n")
	fmt.Scanln(&valor)

	valores := strings.Split(valor, ",")
	x1, _ := strconv.ParseFloat(valores[0], 32)
	y1, _ := strconv.ParseFloat(valores[1], 32)
	p1 := figuras.Punto{X: x1, Y: y1}

	fmt.Print(" Inserte radio... \n")
	fmt.Scanln(&valor)

	radio, _ := strconv.ParseFloat(valor, 32)

	r := figuras.Circulo{Centro: p1, Radio: radio}
	return r
}

func mostrarTexto() string {
	valor := ""
	fmt.Print("Elija su figura...\n- a. Rectangulo \n- b. Cuadrado \n- c. Circulo \n")

	fmt.Scanln(&valor)

	return valor
}
func main() {
	//	fmt.Printf(crearRectangulo().ToString())
	arreglo_figuras := [5]figuras.Figura{}
	valor := mostrarTexto()

	for i := 0; i < 5; i++ {

		if valor == "a" {
			arreglo_figuras[i] = crearRectangulo()
			fmt.Print("- Figura creada -\n")
			valor = mostrarTexto()
		} else if valor == "b" {
			arreglo_figuras[i] = crearCuadrado()
			fmt.Print("- Figura creada -\n")
			valor = mostrarTexto()
		} else if valor == "c" {
			arreglo_figuras[i] = crearCirculo()
			fmt.Print("- Figura creada -\n")
			valor = mostrarTexto()
		} else {
			fmt.Print("Opcion invalida")
			return
		}
	}
	fmt.Print("----------  Figuras guardadas  -----------\n")
	// p1 := figuras.Punto{X: 0, Y: 0}
	// p2 := figuras.Punto{X: 10, Y: 5}
	//
	// r := figuras.Rectangulo{P1: p1, P2: p2}
	//	c := figuras.Cuadrado{Pto: p1, Lado: 10}
	// circulo := figuras.Circulo{Centro: p1, Radio: 5}
	//c := crearCirculo()
	// arreglo_figuras := [5]figuras.Figura{r, circulo}

	for _, f := range arreglo_figuras {
		mostrar(f)
	}
}
