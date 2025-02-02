package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Lulli int    `json:"lulli"`
}

func main() {

	// to convert an object data into json in golang we have to respresent that object as a structure in golang , this is done by struct keyword like in C
	person := Person{"asif", 24, 2}
	fmt.Println("the person is : ", person)

	// ENCODING
	jsonPerson, err := json.Marshal(person) //marshal function is used to encode the data into json format
	if err != nil {
		fmt.Println("the error encountered is : ", err)
		return
	}
	fmt.Println("the json object of person is : ", string(jsonPerson))

	//DECODING
	var personUnjsonified Person
	err = json.Unmarshal(jsonPerson, &personUnjsonified)
	if err != nil {
		fmt.Println("error is : ", err)
		return
	}
	fmt.Println("the person unJSONified is : ", personUnjsonified)

}
