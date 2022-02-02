package config

import (
	"github.com/golang/glog"
	"github.com/jinzhu/configor"
)

type Settings struct {
	Port             string `default:"9090"`
	Debug            bool   `default:"false"`
	SolrInternalHost string `default:"http://localhost:8983"`
	SolrExternalHost string `default:"http://localhost:8983"`
}

var Config = Settings{}

func init() {
	if err := configor.Load(&Config, "config.yml"); err != nil {
		glog.Fatal(err)
	}
}
