package libs

type Square struct {
	SideLength float64
}

func (s Square) getArea() float64 {
	return s.SideLength * s.SideLength
}
