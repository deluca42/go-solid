package shapes

// Rectangle representa um retângulo
type Rectangle struct {
	Largura float64
	Altura  float64
}

// NewRectangle cria uma instância de Rectangle
func NewRectangle(Largura, Altura float64) *Rectangle {
	return &Rectangle{Largura: Largura, Altura: Altura}
}

// Name retorna o nome da forma
func (r *Rectangle) Name() string {
	return "Retangulo"
}

// Area calcula a área do retângulo
func (r *Rectangle) Area() float64 {
	return r.Largura * r.Altura
}
