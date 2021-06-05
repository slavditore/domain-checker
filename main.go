package main

import (
	"log"

	configv1 "gitlab.com/slavditore/check-domain/v2/internal/config.v1"
)

func main() {
	file := "./config/config.yml"
	log.Println("This is a reload of the project")
	log.Println("Trying to read file")

	domains := configv1.ReadConfig(&file)
	log.Printf("%v", domains)
}
