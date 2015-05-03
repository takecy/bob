package entity

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
