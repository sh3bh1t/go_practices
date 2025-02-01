// slices are like vectors in cpp , that is they have dynamic memory allocation and no need to specify size at the time of initialization
// more preferred over arrays
// uses arrays under the hood

package main

import "fmt"

func main() {

	var prices = []int{2, 34, 4, 23} // initialization of slice with values
	fmt.Println(prices)
	prices = append(prices, 69, 78, 30) // adds values to end of slices
	fmt.Println(prices)
	fmt.Printf("the type of prices slice is : %T\n", prices) // prints the data type of prices slice, done by %T
	names := []string{}                                      // to initialize a slice of size 0
	fmt.Println(names)
	fmt.Println(len(names)) // same as arrays to get length of slice

}
