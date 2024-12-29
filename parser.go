package main

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

// //////////////////////////////////////////////////////////////

func ParseYml(filepath string) (*IronAdressbookObj, error) {
	yamlFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	yamlData, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	jsonData, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}

	ironObj := IronAdressbookObj{}
	err = json.Unmarshal(jsonData, &ironObj)
	return &ironObj, err
}

func WriteJson(filepath string, obj *IronAdressbookObj) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(obj); err != nil {
		return err
	}

	return nil
}
