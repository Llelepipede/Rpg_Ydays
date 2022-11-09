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

func Split(str string) []string {
	runes := []rune(str)
	count := 0
	prevIsLeter := false
	for _, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			if prevIsLeter {
				count++
				prevIsLeter = false
			}
		} else {
			prevIsLeter = true
		}
	}
	if prevIsLeter {
		count++
	}
	resArr := make([]string, count)
	i := 0
	start := 0
	prevIsLeter = false
	for ind, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			if prevIsLeter {
				resArr[i] = string(runes[start:ind])
				i++
				prevIsLeter = false
			}
			start = ind + 1
		} else {
			prevIsLeter = true
		}
	}
	if prevIsLeter {
		resArr[i] = string(runes[start:])
	}
	return resArr
}
