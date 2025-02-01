// golang only has for loop and no other bs like while or do-while loop etc
package main

import "fmt"

func main() {

	for i := 5; i < 10; i++ {
		fmt.Println(i)
	}

	// counter := 0
	// for {
	// 	fmt.Println(counter) // eg of a infinite loop
	// 	counter++
	// }

	counters := 0
	for {
		fmt.Println(counters)
		counters++
		if counters == 100 {
			break
		}
	}

	// IMPORTANT
	// 'range' keyword in GO allows to iterate over a string, array , slice etc and provides their index and value as well

	numbers := []int{23, 69, 80, 78, 53}
	for index, value := range numbers {
		fmt.Printf("the index is %d , and value is %d\n", index, value)
	}

	hello := "hello lavdeya"
	for index, charvalue := range hello {
		fmt.Printf("the index of character is %d , and the character value at that index is %c\n", index, charvalue)
	}
}
