// line that is preceded by 'defer' keyword gets executed just before the ending or closing of the main function
// when multiple lines are preceded by 'defer' keyword then it becomes a LIFO structure like stack

package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func main() {

	fmt.Println("kyaa re lavdeya")
	defer fmt.Println("wohoooooooooooo")
	defer fmt.Println("result of addition is : ", add(7, 7, 7))
	fmt.Println("jaa re lavdeya hui hui")

}
