package main

import "fmt"

func main() {

	x := 9
	if x == 10 {
		fmt.Println(x)
	} else if x < 10 {
		fmt.Println("some other value less than 10")
	} else {
		fmt.Println("value greater than 10")
	}
	// rest logical operators go as usual like && , || , == etc , figure it out

}
