package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#784fd0",
		"white": "#987abc",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[int]string)

	colors3[100] = "one hundred"
	fmt.Println(colors3)
	delete(colors3, 100)
	fmt.Println(colors3)

	printMap(colors)

}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("color: ", color, ", hex: ", hex)
	}
}
