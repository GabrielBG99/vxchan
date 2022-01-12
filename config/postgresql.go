package config

type postgresSQL struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	Port     string `yaml:"port"`
	Timezone string `yaml:"timezone"`
	SSL      bool   `yaml:"ssl"`
}
