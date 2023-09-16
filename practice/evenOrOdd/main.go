package main

import "fmt"

func main() {
	var sn []int
	for i := 0; i <= 10; i++ {
		sn = append(sn, i)
	}
	for _, v := range sn {
		if v%2 == 0 {
			fmt.Println(v, "  is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
