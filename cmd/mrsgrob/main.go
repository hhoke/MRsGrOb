package main

import (
	"fmt"
	"log"
	"os"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

type sgrobConf struct {
	Dicom struct {
		Fields []struct {
			Name   string `yaml:"name"`
			Tag    []int  `yaml:"tag"`
			Action struct {
				//going to need to differentiate between nil and empty string
				// use pointer?
				ReplaceWith *string `yaml:"replace-with,omitempty"`
				Delete      *bool   `yaml:"delete,omitempty"`
				NewUID      *bool   `yaml:"new-uid,omitempty"`
			} `yaml:"action"`
		} `yaml:"fields"`
	} `yaml:"dicom"`
}

// loadYamlFile loads a yaml file of the format defined here:
//https://raw.githubusercontent.com/hhoke/mrscrub/main/mrscrub/configs/SSBC_v1.0.yaml
//from a full filename.
func (c *sgrobConf) loadYamlFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	log.Println("foo")
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		return err
	}

	log.Println("bar")
	return nil
}

func (c *sgrobConf) pprintStruct() error {
	//MarshalIndent
	indentedJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("MarshalIndent funnction output %s\n", string(indentedJSON))
	return nil
}

func main() {

	log.Println("loading conf")
	conf := &sgrobConf{}
	fmt.Printf("%#v", conf)
	log.Println("loading conf")
	conf.pprintStruct()
	err := conf.loadYamlFile("/Users/harris.hoke/personal_projects/mrscrub/mrscrub/configs/SSBC_v1.0.yaml")
	if err != nil {
		log.Fatal(err)
	}
	conf.pprintStruct()
}
