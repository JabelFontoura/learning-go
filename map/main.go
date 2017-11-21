package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#F00",
		"green": "#0F0",
		"white": "#FFF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
