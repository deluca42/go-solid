package main

import (
	"fmt"
	"go-solid/shapes"
)

func main() {
	circle := shapes.NewCircle(5)
	printShapeInfo(circle)

	rectangle := shapes.NewRectangle(4, 6)
	printShapeInfo(rectangle)
}

func printShapeInfo(s shapes.Shape) {
	fmt.Printf("Area do %s: %.2f\n", s.Name(), s.Area())
}
