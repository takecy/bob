package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	gojenkins "github.com/yosida95/golang-jenkins"
)

// Bob include jenkins info
type Bob struct {
	JenkinsURLs   map[string]string
	JenkinsAuthes map[string]*gojenkins.Auth

	FilePath string
}

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
