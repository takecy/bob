package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Bob include jenkins info
type Bob struct {
	JenkinsURLs map[string][]ProductConfig

	FilePath string
}

// JenkinsConfig is jenkins config
type JenkinsConfig struct {
	URL   string `yaml:"url"`
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

// EnvironmentConfig is env config
type EnvironmentConfig map[string][]JenkinsConfig

// ProductConfig is product config
type ProductConfig map[string][]EnvironmentConfig

// NewConfig is constructor for Config
func NewConfig(filePath string) (bob *Bob, err error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	bob, _ = NewFromString(string(b))
	bob.FilePath = filePath

	return bob, err
}

// NewFromString is build config from string
func NewFromString(config string) (bob *Bob, err error) {

	bob = new(Bob)
	err = yaml.Unmarshal([]byte(config), &bob)
	if err != nil {
		return nil, err
	}

	return
}
