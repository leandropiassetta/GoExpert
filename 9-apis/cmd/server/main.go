package main

import "github.com/leandropiassetta/goexpert/9-apis/configs"

func main() {
	// config := configs.NewConfig()
	// println(config.GetDbDriver())

	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	println(config.DBDriver)
}
