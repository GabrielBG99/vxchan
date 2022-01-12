package config

type logger struct {
	Name     string `yaml:"name"`
	Level    string `yaml:"level"`
	Filepath string `yaml:"filepath"`
}
