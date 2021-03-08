package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Setting ...
type Setting struct {
	Db struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}
	Listen struct {
		Address string `json:"host"`
		Port    string `json:"port"`
	}
}

// LoadConfiguration ...
func LoadConfiguration(file string) (Setting, error) {
	var setting Setting
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&setting)
	return setting, err
}
