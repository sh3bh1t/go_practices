package main

import (
	"fmt"
	"sync"
)

func sayhello(i int, wg *sync.WaitGroup) {
	defer wg.Done() // before completion for each task in wg , it takes them out of waitgrp , signaling that their execution has been completed
	fmt.Printf("task started %d\n", i)
	fmt.Printf("task completed %d\n", i)
}

func main() {

	var wg sync.WaitGroup //creates a waitgrp
	for i := 1; i <= 3; i++ {
		wg.Add(1)           //increases the watigrp counter by one, that is adds each process to the waitgrp
		go sayhello(i, &wg) // 'go' keyword creates a go routine for the function , i.e it is used to make the function concurrent
	}

	wg.Wait() // waits for all the task in waitgroup to be completed before further execution
	fmt.Println("all tasks completed")

}
