package config

import (
	"encoding/json"
	"io/ioutil"
)

type Configurator struct {
	Ip   string
	Port string
}

func Readconfig(path string) (Configurator, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return Configurator{}, err
	}
	var conf Configurator
	err = json.Unmarshal(f, &conf)
	if err != nil {
		return Configurator{}, err
	}
	return conf, nil
}
