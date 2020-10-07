package config

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

var (
	isLoad = false
	config map[string]interface{}
)

func Load() {
	workDir, _ := os.Getwd()

	file, err := os.Open(workDir + "/config.json")
	if err != nil {
		panic(err)
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
				panic(err)
			}
		}
	}

	toJsonErr := json.Unmarshal([]byte(configString), &config)
	if toJsonErr != nil {
		panic(toJsonErr)
	}
	isLoad = true
}

func Read(key string) interface{} {
	if !isLoad {
		Load()
		isLoad = true
	}
	return config[key]
}