package main

import (
	"flag"
	"net/http"

	"github.com/GabrielBG99/vxchan/config"
	"github.com/GabrielBG99/vxchan/controller/rest"
	"github.com/GabrielBG99/vxchan/logger"
	"github.com/GabrielBG99/vxchan/service"
	"github.com/GabrielBG99/vxchan/storage/fs"
	"github.com/GabrielBG99/vxchan/storage/postgresql"
)

func main() {
	configFile := flag.String("configFile", "config.yaml", "Config File full path. Defaults to current folder")

	flag.Parse()

	if err := config.Load(*configFile); err != nil {
		panic(err)
	}

	log, err := logger.Init(logger.Config{
		Name:     config.Config.Logger.Name,
		Level:    logger.Level(config.Config.Logger.Level),
		Filepath: config.Config.Logger.Filepath,
	})
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	storage, err := postgresql.NewConnector(postgresql.Config{
		Host:     config.Config.Database.Host,
		User:     config.Config.Database.User,
		Password: config.Config.Database.Password,
		DBName:   config.Config.Database.DBName,
		Port:     config.Config.Database.Port,
		Timezone: config.Config.Database.Timezone,
		SSL:      config.Config.Database.SSL,
	})
	if err != nil {
		log.Panic(err)
	}

	fileStorage, err := fs.NewService(config.Config.FileStorage.Folder)
	if err != nil {
		log.Panic(err)
	}

	service := service.NewService(storage, storage, fileStorage)

	if err := service.Seed(); err != nil {
		log.Panic(err)
	}

	router := rest.NewRouter(*service)
	if err := http.ListenAndServe(":5555", router); err != nil {
		log.Fatal(err)
	}
}
