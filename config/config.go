package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	Token string
}

var Config Conf

func ReadConfig() error {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Config)

	if err != nil {
		return err
	}
	fmt.Printf("Config.Token: %v\n", Config.Token)
	return nil
}
