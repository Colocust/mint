package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
)

var (
	isLoad = false
	config map[string]interface{}
)

func Load(name string) error {
	if isLoad == false {
		file, err := os.Open(name)
		if err != nil {
			return errors.New("can not find " + name)
		}

		var configString string

		buf := bufio.NewReader(file)
		for {
			item, err := buf.ReadString('\n')
			configString += item
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
		}
		if err = file.Close(); err != nil {
			return err
		}

		if err = json.Unmarshal([]byte(configString), &config); err != nil {
			return err
		}

		isLoad = true
	}
	return nil
}

func Get(key string) interface{} {
	return config[key]
}
