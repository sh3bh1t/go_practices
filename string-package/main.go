package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,mango,orange"
	data_parts := strings.Split(data, ",") //seperates values based on seperator provided as second parameter and returns those values in an array
	fmt.Println(data_parts)
	data1 := "apple, banana, mango, orange"
	data_parts1 := strings.Split(data1, ", ")
	fmt.Println(data_parts1)

	str := "lala , mallu ,mulla, kattua ,mulla, kattua"
	count := strings.Count(str, "mulla") // returns count of no of times a word appears
	fmt.Println(count)

	str1 := "                  beteeeee, mai hu lucky the racer  ! ! !       "
	trimmed := strings.TrimSpace(str1) // trims the leading and trailing whitespaces within a string
	fmt.Println(trimmed)

	strjoined := strings.Join([]string{"nirmala", "tai", "haye", "haye"}, " ") // takes an array of strings as input as first parameter and concatenates them with seperator provided as 2nd param b/w them
	fmt.Println(strjoined)
}
