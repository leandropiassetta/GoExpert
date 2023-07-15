package main

import (
	"fmt"
	"strings"
)

func throwMessageError(errMsg string) string {
	return fmt.Sprint("return other message")
}

func registerMessage(errMsg string) {
	fmt.Println("this message was registered: ", errMsg)

	// return fmt.Sprintf("this message was registered: ", errMsg)
}

func testUseCase(err string) string {
	var msg string

	msg = "this message is: not found"
	msg2 := "this message is: found"

	condicionA := strings.Contains(msg, err)
	condicionB := strings.Contains(msg2, err)

	switch {
	case condicionA:
		msg = throwMessageError(msg)
	case condicionB:
		msg = throwMessageError(msg)
		// println("condicionB", msg)
	default:
		registerMessage(msg)
	}

	return msg
}

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

	fmt.Println(testUseCase("not found"))
}
