package main

import (
	"log"

	certcheck "gitlab.com/slavditore/check-domain/v2/internal/cert_check"
	configv1 "gitlab.com/slavditore/check-domain/v2/internal/config.v1"
	output_plain "gitlab.com/slavditore/check-domain/v2/internal/output_log"
)

func main() {
	file := "./config/config.yml"
	log.Println("This is a reload of the project")
	log.Println("Trying to read file")
	domains := configv1.ReadConfig(&file)
	log.Println("Process the logs")
	domains = certcheck.CheckDomains(domains)
	output_plain.PlainOutput(&domains)
	log.Println("Application is finished")
}
