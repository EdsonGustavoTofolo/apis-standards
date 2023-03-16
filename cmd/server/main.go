package main

import "github.com/EdsonGustavoTofolo/apis-standards/configs"

func main() {
	cf := configs.LoadConfig("./cmd/server/")
	println(cf.DBDriver)
}
