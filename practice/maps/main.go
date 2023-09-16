package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	//var colors map[string]string
	//colors := make(map[string]string)

	colors["white"] = "#ffffff"

	printMap(colors)

	delete(colors, "white")

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("color: %v hexcode: %v\n", color, hex)
	}
}
