package config

import (
	"github.com/golang/glog"
	"github.com/jinzhu/configor"
)

type Settings struct {
	Port       string `default:"9090"`
	Debug      bool   `default:"false"`
	SearchHost string `default:"http://localhost:8983"`
}

var Config = Settings{}

func init() {
	if err := configor.Load(&Config, "config.yml"); err != nil {
		glog.Fatal(err)
	}
}
