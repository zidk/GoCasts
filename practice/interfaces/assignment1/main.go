package main

import "assignment1/libs"

func main() {
	s := libs.Square{SideLength: 10}
	t := libs.Triangle{
		Height: 10,
		Base:   5,
	}
	libs.PrintArea(s)
	libs.PrintArea(t)
}
