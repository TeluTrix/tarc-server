package service

import (
	"log/slog"
	"os"

	"github.com/TeluTrix/tarc-server/internal/model"
	"github.com/pelletier/go-toml/v2"
)

var Conf model.Config

func ReadConfig() {
	dat, err := os.ReadFile("/etc/tarc/conf.toml")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	errO := toml.Unmarshal([]byte(dat), &Conf)
	if errO != nil {
		slog.Error(errO.Error())
		panic(errO)
	}
}
