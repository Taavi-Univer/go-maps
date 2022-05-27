package main

import "fmt"

func main() {
	colors := map[string]string{ // all keys are of type string and all values are of type string
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// var colors map[string]string
	// colors["white"] = "#ffffff"

	// colors := make(map[string]string)

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
