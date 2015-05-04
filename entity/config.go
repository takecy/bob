package entity

// Bob include jenkins info
type Bob struct {
	ProductConfig ProductConfig
	FilePath      string
}

// ProductConfig key:prodcutName value:EnvironmentConfig
type ProductConfig map[string]EnvironmentConfig

// EnvironmentConfig key:environment value:JenkinsConfig
type EnvironmentConfig map[string]JenkinsConfig

// JenkinsConfig is jenkins information
type JenkinsConfig struct {
	URL   string `yaml:"url"`
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}
