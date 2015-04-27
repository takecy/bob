package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Bob include jenkins info
type Bob struct {
	ProductConfig *ProductConfig
	FilePath      string
}

// ProductConfig is product config
type ProductConfig map[string]EnvironmentConfig

// EnvironmentConfig is env config
type EnvironmentConfig map[string]JenkinsConfig

// JenkinsConfig is jenkins config
type JenkinsConfig struct {
	URL   string `yaml:"url"`
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

// NewConfig is constructor for Config
func NewConfig(filePath string) (bob *Bob, err error) {
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
func NewFromString(config string) (bob *Bob, err error) {

	yamlConfig := make(ProductConfig, 0)
	err = yaml.Unmarshal([]byte(config), &yamlConfig)
	if err != nil {
		return nil, err
	}

	bob = new(Bob)
	bob.ProductConfig = &yamlConfig

	return
}
