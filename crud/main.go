package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// struct to store the data from response of GET method
type Todo struct {
	UserId    int    `json="u_id"`
	Id        int    `json="id"`
	Title     string `json="Title`
	Completed bool   `json="status"`
}

func performGetRequest() {
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

func performPostRequest() {

	todo := Todo{
		UserId:    24,
		Title:     "tmkc",
		Completed: true,
	}

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	//convert todo data into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshaling : ", err)
		return
	}

	// convert the json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	// sending POST request
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error sending req : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
	fmt.Println("response status: ", res.Status)

}

func performUpdateRequest() {
	todo := Todo{
		UserId:    2989,
		Title:     "tmkcOnceAgain",
		Completed: false,
	}

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	//convert todo data into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshaling : ", err)
		return
	}

	//create PUT request
	//GO library of http only has GET and POST methods so for PUT method we need to create a new request and that function takes three parameters in which the third one is a reader form

	// convert the json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	//create PUT req
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("error putting : ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request : ", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
	fmt.Println("response status  for putting : ", res.Status)
}

func performDeleteRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	//create DELETE req
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("error deleting : ", err)
		return
	}
	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request : ", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
	fmt.Println("response status  for putting : ", res.Status)
}

func main() {
	fmt.Println("learning CRUD IN GO...... ")

	//performGetRequest()
	//performPostRequest()
	//performUpdateRequest()
	performDeleteRequest()
}
