package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
	WaitSeconds     int           // Period of the timer checks in seconds
	StartupCheck    bool          // Check miners on startup
	Log             bool          //Enable or disable logging
	TgBotActivate   bool          //Enable or disable Telegram bot
	TgAPIKey        string        //Telegram Api key for bot communicationg
	TgAdminUserName string        //Telegram Username which will control the bot
	Pushover        bool          //Enable or disable Pushover notifications
	PushoverToken   string        //Pushover access token
	PushoverUser    string        //Pushover user token
	Miners          []MinerConfig // An array of the
}

//Config is the global Config variable
var Config ConfigurationFile

//ReadConfig - read and parse the config file
func ReadConfig() (configFile ConfigurationFile) {
	//get binary dir
	//os.Args doesn't work the way we want with "go run". You can use next line
	//for local dev, but use the original for production.
	dir, err := filepath.Abs("./")
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	log.Notice("Reading file config.json...")
	file := dir + "/config.json"

	configFileContent, err := ioutil.ReadFile(file)
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

	if configFile.Pushover == true {
		log.Notice("Pushover notification is ENABLED")
		log.Notice("Pushover Token:", configFile.PushoverToken)
		log.Notice("Pushover User:", configFile.PushoverUser)
	}

	return
}
