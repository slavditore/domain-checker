package main

import (
	"log"

	flag "github.com/spf13/pflag"
	certcheck "gitlab.com/slavditore/check-domain/v2/internal/cert_check"
	configv1 "gitlab.com/slavditore/check-domain/v2/internal/config.v1"
	output_plain "gitlab.com/slavditore/check-domain/v2/internal/output_log"
)

var (
	argFileName = flag.String("config", "./config/config.yml", "path to file with domains")
)

func main() {
	flag.Parse()
	log.Println("This is a reload of the project")
	log.Println("Trying to read file")
	domains, err := configv1.ReadConfig(argFileName)
	if err != nil {
		log.Printf("unable to process the file: %v", err)
		return
	}
	log.Println("Process the logs")
	domains = certcheck.CheckDomains(domains)
	output_plain.PlainOutput(&domains)
	log.Println("Application is finished")
}
