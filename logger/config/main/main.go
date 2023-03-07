package main

import (
	"fmt"

	"service/logger/service/config"
)

func main() {
	fmt.Println(config.DBConfig())
	fmt.Println(config.DBCollection())
}
