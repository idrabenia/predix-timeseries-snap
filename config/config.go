package config

import (
	"io/ioutil"
	"os"
	"time"
	"gopkg.in/yaml.v2"
)

type TsConfig struct {
	IngestUrl string
	ZoneId string
	UaaIssuerUrl string
	ClientId string
	ClientSecret string
}

func Load(fileName string) TsConfig {
	filePath := configPath(fileName)

	waitOn(filePath)

	bytes, err := ioutil.ReadFile(filePath)
	panicOnError(err)

	config := TsConfig{}
	err = yaml.Unmarshal(bytes, &config)
	panicOnError(err)

	return config
}

func (config *TsConfig) Save(fileName string) {
	str, err := yaml.Marshal(config)
	panicOnError(err)

	filePath := configPath(fileName)
	writeErr := ioutil.WriteFile(filePath, []byte(str), 0644)
	panicOnError(writeErr)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func configPath(fileName string) string {
	snapCommonDir, isDirectorySet := os.LookupEnv("SNAP_COMMON")

	if !isDirectorySet {
		snapCommonDir = "."
	}

	return snapCommonDir + "/" + fileName
}

func waitOn(filePath string) {
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// nope
		} else {
			return
		}

		time.Sleep(300)
	}
}

