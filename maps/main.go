package main

import "fmt"

func main() {

	// name <-> marks  , here name is the KEY of string datatype and marks is VALUE of integer datatype
	students_marks := make(map[string]int) // used to initialize map in GO

	students_marks["jhaatu"] = 89 // key is 'jhaatu' and value set for it is 89
	students_marks["kattua"] = 00
	students_marks["chootpaglu"] = 78

	fmt.Println("the marks of kattua is : ", students_marks["kattua"]) // accessing the value of key 'kattua'

	delete(students_marks, "chootpaglu") //delete all data for key 'chootpaglu

	students_marks["jhaatu"] = 96 //updating value for 'jhaatu' key
	fmt.Println("marks of jhaatu is : ", students_marks["jhaatu"])

	_, exists := students_marks["jhaatu"] // accessing value for a key in map always returns two thing the VALUE of the KEY , and BOOLEAN VALUE that the key exists or not, true if exists and vice-versa
	println("jhaatu exists : ", exists)

	for index, value := range students_marks { //for iterating over the map
		fmt.Printf("the key is %s , marks is %d\n", index, value)
	}
}
