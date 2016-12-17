package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ServerConfig   ServerConfig
	StorageConfig  StorageConfig
}

type ServerConfig struct {
	Port string
}

type StorageConfig struct {
	ProviderType string
	FileStorage FileStorageConfig
}

type FileStorageConfig struct {
	RootPath string
}

func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
	if o.StorageConfig.ProviderType == "" {
		o.StorageConfig.ProviderType = "File"
	}
	if o.StorageConfig.FileStorage.RootPath == "" {
		o.StorageConfig.FileStorage.RootPath = "./repository/"
	}
}

var ConfigParam *Config = &Config{}

func findConfigFile(fileName string) string {
	if _, error := os.Stat("./" + fileName); error == nil {
		fileName, _ = filepath.Abs("./" + fileName)
	} else if _, error := os.Stat("./config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, error := os.Stat("./src/github.com/MiteshSharma/StorageService/config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./src/github.com/MiteshSharma/StorageService/config/" + fileName)
	}
	return fileName
}

func LoadConfig(fileName string) {
	filePath := findConfigFile(fileName)

	file, error := os.Open(filePath)

	if error != nil {
		panic("Error occured during config file reading " + error.Error())
	}

	jsonParser := json.NewDecoder(file)

	config := Config{}

	if jsonErr := jsonParser.Decode(&config); jsonErr != nil {
		panic("Json parsing error" + jsonErr.Error())
	}

	config.SaveDefaultConfigParams()

	ConfigParam = &config
}
