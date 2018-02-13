package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config struct of json config file
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

// New return config struct read values from config json file
func New(file string) (Config, error) {
	c := Config{}

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(dat, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
