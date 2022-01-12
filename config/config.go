package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Config *config

type config struct {
	Database    postgresSQL `yaml:"postgres"`
	FileStorage fileStorage `yaml:"fileStorage"`
	Logger      logger      `yaml:"logger"`
}

func Load(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	Config = new(config)
	return yaml.NewDecoder(f).Decode(Config)
}
