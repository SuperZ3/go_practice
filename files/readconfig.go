package main

import (
	"encoding/json"
	"os"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		nameSlice := make([]byte, 5)
		file.ReadAt(nameSlice, 20)
		Config.UserName = string(nameSlice)
		file.Seek(48, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config.AdditionalProducts)
	}
	return
}
func init() {
	err := LoadConfig()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	}
}
