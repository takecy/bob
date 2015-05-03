package config

import (
	"io/ioutil"

	"github.com/takecy/bob/entity"
	"gopkg.in/yaml.v2"
)

// NewConfig is constructor for Config
func NewConfig(filePath string) (bob *entity.Bob, err error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	bob, err = NewFromString(string(b))
	if err != nil {
		return nil, err
	}

	bob.FilePath = filePath

	return
}

// NewFromString is build config from string
func NewFromString(config string) (bob *entity.Bob, err error) {

	yamlConfig := make(entity.ProductConfig, 0)
	err = yaml.Unmarshal([]byte(config), &yamlConfig)
	if err != nil {
		return nil, err
	}

	bob = new(entity.Bob)
	bob.ProductConfig = &yamlConfig

	return
}
