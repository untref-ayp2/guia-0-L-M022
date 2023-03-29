package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Centro Punto
	Radio  float64
}

func (c Circulo) Area() float64 {

	return math.Pi * c.Radio * c.Radio
}
func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circulo) ToString() string {
	_string := fmt.Sprint("Circulo:", c)

	return _string
}
