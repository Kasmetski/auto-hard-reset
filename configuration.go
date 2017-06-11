package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//MinerConfig struct for json parse
type MinerConfig struct {
	Name string // NAME
	Pin  string // PIN-NUMBER OF GPIO
	IP   string // IP ADDRESS
	Info string // ADDITIONAL INFO
}

//ConfigurationFile struct to parse config.json
type ConfigurationFile struct {
	WaitSeconds int           // Period of the timer checks in seconds
	Miners      []MinerConfig // An array of the
}

//ReadConfig - read and parse the config file
func ReadConfig() (configFile ConfigurationFile) {
	log.Notice("Reading file config.json...")
	configFileContent, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Error("Trying to read file config.json, but:", err)
		os.Exit(1)
	}

	log.Notice("Parsing configuration file...")
	err = json.Unmarshal(configFileContent, &configFile)
	if err != nil {
		log.Error("Parsing JSON content, but:", err)
		os.Exit(2)
	}

	log.Notice("Timer (time in seconds):", configFile.WaitSeconds)
	log.Notice("Found miner configurations:", len(configFile.Miners))

	return
}
