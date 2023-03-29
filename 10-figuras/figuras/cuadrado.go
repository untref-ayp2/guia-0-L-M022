package figuras

import "fmt"

type Cuadrado struct {
	Pto  Punto //esquina inferior izquierda
	Lado float64
}

func (c *Cuadrado) Mover(incX, incY float64) {
	c.Pto.Mover(incX, incY)
}

func (c Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) ToString() string {
	cadena := fmt.Sprint("Cuadrado:", c)

	return cadena
}
