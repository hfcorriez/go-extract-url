package main

import (
	"github.com/callum-ramage/jsonconfig"
)

var config jsonconfig.Configuration

func InitConfig () {
	var err error
	config, err = jsonconfig.LoadAbstract("./config/default.json", "")

	if err != nil {
		return
	}
}

func Config(key string) jsonconfig.JSONValue {
	return config[key]
}