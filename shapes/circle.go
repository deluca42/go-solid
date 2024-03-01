package shapes

import "math"

// Circle representa um círculo
type Circle struct {
	Raio float64
}

// NewCircle cria uma instância de Circle
func NewCircle(Raio float64) *Circle {
	return &Circle{Raio: Raio}
}

// Name retorna o nome da forma
func (c *Circle) Name() string {
	return "Circulo"
}

// Area calcula a área do círculo
func (c *Circle) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
