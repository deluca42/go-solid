package shapes

// Shape é uma interface para formas geométricas
type Shape interface {
	Name() string
	Area() float64
}