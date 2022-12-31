package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println("A:", a)
	} else {
		println("B:", b)
	} // in Go not exist "else if"

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b || c > a")
	}

	switch a + b {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("none of those")
	}
}
