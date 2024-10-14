package app

import (
	"encoding/json"
	"fmt"
	"os"
)

func BlankConfig() Config {
	return Config{Sites: []SiteInfo{}}
}

type Config struct {
	Email string `json:"email"`
	ApiKey string `json:"apikey"`
	CurrentIp string `json:"currentip"`
	Sites Sites `json:"sites"`
}


type SiteInfo struct {
	ZoneId string `json:"zoneid"`
	RecordId string `json:"recordid"`
	URI string `json:"uri"`
	BaseURI string `json:"baseuri"`
	IP string `json:"ip"`
}

type Sites []SiteInfo

func (c Config) SaveConfigFile(path string) {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal the json while saving the config file.\nGot error %s\n", err)
		os.Exit(2)
	}
	err = os.WriteFile(path, b, 0644)
	if err !=nil {
		fmt.Printf("Failed to save config file.\nGot error %s\n",err)
		os.Exit(2)
	}
}

func LoadConfigFile(path string) Config {
	b, err := os.ReadFile(path)
	if err != nil{
		fmt.Printf("Failed to load config file from %s.\nGot error %s\n", path, err)
		os.Exit(2)
	}	
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil{
		fmt.Printf("Failed to parse json data while loading config file.\nGot error %s\n", err)	
		os.Exit(2)
	}	
	return config
}

func (sites Sites) FindSiteIndex(uri string) int{
	for i, s := range sites {
		if s.URI == uri {
			return i
		}
	}	
	return -1
}

func (sites Sites) Exists(uri string) bool{
	for _, s := range sites {
		if s.URI == uri {
			return true 
		}
	}	
	return false 
}