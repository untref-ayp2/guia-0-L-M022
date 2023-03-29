package figuras

import "fmt"

type Punto struct {
	X float64
	Y float64
}

func (p *Punto) ToString() string {
	cadena := fmt.Sprintf("(%v, %v)", p.X, p.Y)
	return cadena
}

func (p *Punto) Mover(x, y float64) {
	p.X += x
	p.Y += y
}
