package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
	}
	fmt.Println(colors)

	printMap(colors)

	colors2:=make(map[string]string)

	colors2["white"]="#ffffff" // Add new key

	fmt.Println(colors2)


	delete(colors,"red")

	fmt.Println(colors)
}


func printMap(c map[string]string){
	for color ,hex:=range c{
		fmt.Println("Hex code for", color,"is",hex)
	}
}