package postgresql

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	Timezone string
	SSL      bool
}

func (c Config) DSN() string {
	sslStatus := "disable"
	if c.SSL {
		sslStatus = "enable"
	}

	dsnStr := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	return fmt.Sprintf(dsnStr, c.Host, c.User, c.Password, c.DBName, c.Port, sslStatus, c.Timezone)
}
