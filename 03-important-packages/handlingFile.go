package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creating file
	f, err := os.Create("archive.txt")
	if err != nil {
		panic(err)
	}

	// this save strings
	// lenght, err := f.WriteString("Hello, World!")

	// if know that the i will go save is a string, i utilize WriteString()
	// if not i utilize Write()

	// this save bytes
	lenght, err := f.Write([]byte("Writing data in the file, for this experience!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("file created with success! lenght: %d bytes", lenght)
	f.Close()

	// Reading file

	file, err := os.ReadFile("archive.txt")
	if err != nil {
		panic(err)
	}

	// convert in string because always it is registered in the file, it is registered as slice of bytes
	fmt.Println("\n", string(file))

	// Atention!
	// How i read a file that is more bigger that the memory in my computer??
	// In Go i get this, reading piece for piece this file

	// in this cases! we read little by little

	file2, err := os.Open("archive.txt")
	if err != nil {
		panic(err)
	}

	// In this we case, i created a Reader -> Reader have a capacity for read this content, but this content will go be buffered, that is, it will be read gradually

	reader := bufio.NewReader(file2)

	// i define how much bytes this reader will read in the file, that is my buffer
	buffer := make([]byte, 4)

	// this loop, will be read this buffer
	for {
		// position = the position where that have doing the reading.
		position, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// will make print for me, converted in string, he is reading this buffer and it's joining in a slice the content that he is getting
		fmt.Println(string(buffer[:position]))
	}

	// removing file

	err = os.Remove("archive.txt")
	if err != nil {
		panic(err)
	}
}
