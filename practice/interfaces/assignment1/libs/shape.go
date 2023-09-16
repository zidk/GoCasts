package libs

import "fmt"

type Shape interface {
	getArea() float64
}

func PrintArea(s Shape) {
	a := s.getArea()
	fmt.Println("Area: ", a)
}
