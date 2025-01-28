package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("biyaah hogya tumhara?")
	//var name string
	//fmt.Scan(&name) //scan method only reads till first whitespace character

	//fmt.Println("\nkoi chhua nhi hai kya ," + name + " ?")

	reader := bufio.NewReader(os.Stdin) // scan only reads till first whitespace char so instead we need to use 'bufio' package to read a whole line and os package for taking input through standard method 'stdin'
	name, _ := reader.ReadString('\n')
	fmt.Println("aaj to balatkaar ki raat h " + name)

}
