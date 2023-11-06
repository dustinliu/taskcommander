package controller

import (
	"strings"

	"github.com/dustinliu/taskcommander/service"
)

var config *Config

const (
	dataLocation = "data.location"
)

func init() {
	config = &Config{}
	data, err := service.Taskwarrior("show")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, dataLocation) {
			str, _ := strings.CutPrefix(line, dataLocation)
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
