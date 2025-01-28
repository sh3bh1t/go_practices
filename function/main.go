package main

import "fmt"

func simpFunction() {
	fmt.Println("wohooooooo")
}

func add(a, b int) int { //here a and b are integer paramaters and int is the return type
	return a + b
}

func main() {
	fmt.Println("today learning about functions in go")
	simpFunction()
	fmt.Println(add(69, 69))
}
