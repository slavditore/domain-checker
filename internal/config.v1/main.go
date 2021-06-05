package configv1

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func (config *ConfigV1) parseConfiguration(filename *string) *ConfigV1 {
	sourceFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Printf("Unable to read file %s: %v", *filename, err)
	}
	err = yaml.Unmarshal(sourceFile, config)
	if err != nil {
		log.Fatalf("Unable to parse file %s: %v", *filename, err)
	}
	return config
}

func ReadConfig(filename *string) {
	var config ConfigV1
	config.parseConfiguration(filename)
	fmt.Println(config)
}
