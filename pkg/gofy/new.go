package gofy

import (
	"github.com/varun-singhh/gofy/pkg/gofy/config"
	"github.com/varun-singhh/gofy/pkg/gofy/log"
	"github.com/varun-singhh/gofy/pkg/gofy/server"
	"os"
	"strconv"
)

func New() (k *Gofy) {
	logger := log.NewLogger()

	return NewWithConfig(config.NewGoDotEnvProvider(logger, getConfigFolder()))
}

func NewWithConfig(c config.Config) (k *Gofy) {
	// Here we do things based on what is provided by Config
	logger := log.NewLogger()

	gofy := &Gofy{
		Logger: logger,
		Config: c,
	}

	s := server.NewServer()
	gofy.Server = s

	// HTTP PORT
	p, err := strconv.Atoi(c.Get("HTTP_PORT"))
	s.HTTP.Port = p

	if err != nil || p <= 0 {
		s.HTTP.Port = 8000
	}

	return gofy
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
