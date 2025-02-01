package main

import "fmt"

func main() {

	var name [5]int //for initializing and by default values set to 0 for int , "" for string and false for boolean
	fmt.Println(name)

	var price = [5]string{"jhaat", "kaattua"} // for initializing and declaring value at same time
	fmt.Println(price)

	prices := [5]string{"jhaat", "kaattua"} //shorthand for the same above
	fmt.Println(prices)

	fmt.Println(name[3]) // for accessing indices

	fmt.Println(len(name)) // for getting the length of array
}
