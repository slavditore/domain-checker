package configv1

import (
	"io/ioutil"
	"log"

	. "gitlab.com/slavditore/check-domain/v2/internal/types"
	"gopkg.in/yaml.v2"
)

func (config *ConfigV1) parseConfiguration(filename *string) *ConfigV1 {
	sourceFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", *filename, err)
	}
	err = yaml.Unmarshal(sourceFile, config)
	if err != nil {
		log.Fatalf("Unable to parse file %s: %v", *filename, err)
	}
	return config
}

func ReadConfig(filename *string) []DomainStatus {
	var config ConfigV1
	var domains []DomainStatus
	config.parseConfiguration(filename)
	for i := 0; i < len(config.Groups); i++ {
		for j := 0; j < len(config.Groups[i].Domains); j++ {
			domains = append(domains, DomainStatus{Name: config.Groups[i].Domains[j], Group: config.Groups[i].Name})
			//fmt.Println(config.Groups[i].Domains[j])
		}
	}
	return domains
}
