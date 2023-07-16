package main

import "github.com/leandropiassetta/goexpert/7-packaging/1/math"

// workspace location, in this case is: github.com/leandropiassetta/goexpert/7-Packaging/4

// workspaces local packages dont prejuidicate the global packages, ocasionaly the local package dont need publish in github.com repository and can be used in other projects in my local machine without need publish in github.com repository and without need use the command go mod edit -replace github.com/leandropiassetta/goexpert/7-Packaging/4/math=../math

// go work init -> create a new workspace in my local machine and i will to write what are  the libraries that i will use in my project, this command will create a file called go.work in my project root folder and this file will have the libraries that i will use in my project, this file is like a package.json in nodejs, this file go.work will let me know version of go that i will use in my project and the libraries that i will use in my project and the relative url of this libraries in my local machine

// eventually i can put this file in the .gitignore file because this file is not necessary in my project, this file is necessary only in my local machine

// go mod tidy -e ->  will ignore the packages that my sisytem dont found in my local machine

func main() {
	math := math.NewMath(1, 2, "mathematic")

	println(math.Sum())
}
