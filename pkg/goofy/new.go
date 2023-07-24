package goofy

import (
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/varun-singhh/gofy/pkg/goofy/config"
	"github.com/varun-singhh/gofy/pkg/goofy/log"
)

func New() (k *Goofy) {
	logger := log.NewLogger()

	return NewWithConfig(config.NewGoDotEnvProvider(logger, getConfigFolder()))
}

func NewWithConfig(c config.Config) (k *Goofy) {
	// Here we do things based on what is provided by Config
	logger := log.NewLogger()

	db := connectPostgresDB(c, logger)

	goofy := &Goofy{
		Logger:   logger,
		Config:   c,
		Database: db,
	}

	s := NewServer(goofy)
	goofy.Server = s

	// HTTP PORT
	p, err := strconv.Atoi(c.Get("HTTP_PORT"))
	s.HTTP.Port = p

	if err != nil || p <= 0 {
		s.HTTP.Port = 8000
	}

	return goofy
}

func getConfigFolder() (configFolder string) {
	if _, err := os.Stat("./configs"); err == nil {
		configFolder = "./configs"
	} else if _, err = os.Stat("../configs"); err == nil {
		configFolder = "../configs"
	} else {
		configFolder = "../../configs"
	}

	return
}

func connectPostgresDB(c config.Config, logger log.Logger) *gorm.DB {
	host := c.Get("DB_HOST")
	name := c.Get("DB_NAME")
	pass := c.Get("DB_PASSWORD")
	root := c.Get("DB_ROOT")
	port := c.Get("DAB_PORT")

	dsn := "host=" + host + " user=" + root + " password=" + pass + " dbname=" + name + " port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error while connecting to DB, Error is %v", err)

		return nil
	}

	logger.Info("DB Connected Successfully")

	return db
}
