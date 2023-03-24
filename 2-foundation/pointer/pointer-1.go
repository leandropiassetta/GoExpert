package main

func soma(a, b *int) int {
	// Have here the copy value of param A, and i can to alter this value
	// For to alter this value in the memory i to use the pointer
	*a = 50
	return *a + *b
}

func main() {
	// Every time that will pass a value, i´m sending a copy this value that in memory

	num1 := 10
	num2 := 20

	println(num1)
	println(soma(&num1, &num2))
}

// When i use the pointer and when i dont use?

// Don't use -> When i want to send a copy of the datas to make a use

// Use -> When for some reason i need use values mutables
