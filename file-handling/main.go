package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("woohoo.txt") // to create a file using OS package t, gives two outputs , the file object created and error
	// if err != nil {
	// 	fmt.Println("the error encountered during file creation is : ", err)
	// }

	// str := "making progress \n wooooohooooooooo lesssssssgoooooooooo"
	// io.WriteString(file, str) // input/output package used for writing streams of data into the file or reading it

	// defer file.Close()

	/*
		file, err := os.Open("woohoo.txt")  //opens a file
		if err != nil {
			fmt.Println("the error encountered during file creation is : ", err)
		}
		defer file.Close()
		buffer := make([]byte, 1024)  // creates a buffer of size 1024 bytes
		for {
			n, err := file.Read(buffer)  // puts content of file into buffer
			if err == io.EOF { // if End Of File reached
				fmt.Println("end of file reached ")
				break
			}
			if err != nil {
				fmt.Println("error encounterd : ", err)
				return
			}

			fmt.Println(string(buffer[:n]))  // converts contents of buffer into string and prints
		}
	*/

	// shortcut method of reading the file
	content, err := os.ReadFile("woohoo.txt") //Readfile reads the entire file and puts the entire file into memory all at once , so a memory issue might be encountered if file size is too big
	if err != nil {
		fmt.Println("error encountered", err)
		return
	}
	fmt.Println(string(content)) // stringifies the contents of file and prints them

}
