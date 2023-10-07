package main

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	print(n, " * ", n-1, "\n")

	return n * factorial(n-1)
}

func main() {
	println(factorial(5))
}
