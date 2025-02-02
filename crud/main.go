package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct to store the data from response of GET method
type Todo struct {
	UserId    int    `json="u_id"`
	Id        int    `json="id"`
	Title     string `json="Title`
	Completed bool   `json="status"`
}

func main() {
	fmt.Println("learning CRUD IN GO...... ")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("the error encounterd is : ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error in getting response :", res.Status)
		return
	}

	// TO READ THE DATA FROM RESPONSE OF AN REQUEST

	// NOT PREFERRED METHOD

	// data, err := ioutil.ReadAll(res.Body)  // BUT WE DO NOT PREFER THIS METHOD TO GET DATA
	// if err != nil {
	// 	fmt.Println("the error encounterd is : ", err)
	// 	return
	// }
	// fmt.Println("the data received from GET is :", string(data))

	// PREFERRED METHOD
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("the error encounterd is : ", err)
		return
	}
	fmt.Println("TODO : ", todo)

}
