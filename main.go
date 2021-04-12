package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./config/config.yml"
	fmt.Println("This is a reload of the project")
	fmt.Println("Trying to read file")
	config, _ := ioutil.ReadFile(file)
	fmt.Println(string(config))
}
