package controller

import (
	"strings"

	"github.com/dustinliu/taskcommander/service"
)

var config *Config

const (
	dataLocationKey = "data.location"
)

func init() {
	config = &Config{}
	data, err := service.Taskwarrior("show")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, dataLocationKey) {
			str, _ := strings.CutPrefix(line, dataLocationKey)
			config.Data_location = strings.TrimSpace(str)
		}
	}
}

type Config struct {
	Data_location string
}

func GetConfig() *Config {
	return config
}
