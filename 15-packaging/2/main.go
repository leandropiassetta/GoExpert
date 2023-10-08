package main

import "github.com/google/uuid"

// go mod tidy -> to download the dependencies, he see the package that we are using and download the dependencies and remove the dependencies that we are not using
func main() {
	println(uuid.New().String())
}
