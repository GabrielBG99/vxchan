package postgresql

import (
	"regexp"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	alreadyExistsErrorPattern = regexp.MustCompile(`Key \(.+\)=\(.+\) already exists\.`)
)

type connector struct {
	db *gorm.DB
}

func (c connector) init() error {
	if err := c.initBoard(); err != nil {
		return err
	}

	if err := c.initThread(); err != nil {
		return err
	}

	return nil
}

func NewConnector(config Config) (*connector, error) {
	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	c := &connector{
		db: db,
	}

	return c, c.init()
}
