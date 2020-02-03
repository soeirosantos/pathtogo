package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServiceConfig struct {
	DbUrlConnection     string `json:"db_url_connection"`
	NotificationChannel string `json:"notification_channel"`
	WeatherSvcEndpoint  string `json:"weather_endpoint"`
}

func main() {
	c, err := loadConfig("./config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(c.DbUrlConnection)
}

func loadConfig(configPath string) (ServiceConfig, error) {
	configFile, err := os.Open(configPath)
	c := ServiceConfig{}
	if err != nil {
		return c, err
	}

	err = json.NewDecoder(configFile).Decode(&c)
	return c, err
}
