package main

import (
	"time"

	logging "github.com/op/go-logging"

	"encoding/json"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"io/ioutil"
	"os"
)

type MinerConfig struct {
	Name string // NAME
	Pin  string // PIN-NUMBER OF GPIO
	Ip   string // IP ADDRESS
	Info string // ADDITIONAL INFO
}

type ConfigurationFile struct {
	WaitSeconds int           // Period of the timer checks in seconds
	Miners      []MinerConfig // An array of the
}

var log = logging.MustGetLogger("auto-hard-reset-log")

func main() {
	log.Notice("Reading file config.json...")
	configFileContent, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Error("Trying to read file config.json, but:", err)
		os.Exit(1)
	}

	log.Notice("Parsing configuration file...")
	var configFile ConfigurationFile
	err = json.Unmarshal(configFileContent, &configFile)
	if err != nil {
		log.Error("Parsing JSON content, but:", err)
		os.Exit(2)
	}

	log.Notice("Waiting time in seconds for the timer:", configFile.WaitSeconds)

	minerConfigs := configFile.Miners
	totalMinerConfigs := len(minerConfigs)
	log.Notice("Found miner configurations:", totalMinerConfigs)

	r := raspi.NewAdaptor()
	///MINING RIGS CONFIGURATION///
	miningRigs := make([]Rig, 0)
	for _, m := range minerConfigs {
		log.Notice("minerConfig:", m)
		miningRigs = append(miningRigs, Rig{m.Name, gpio.NewRelayDriver(r, m.Pin), m.Ip, m.Info})
	}
	log.Notice("Configured rigs: ", len(miningRigs))

	LogMachines()

	check := func() {
		log.Notice("Checking machines: ")
		for i := 0; i < len(miningRigs); i++ {
			log.Notice("Ping miner: ", i, "name: ", miningRigs[i].name, "ip: ", miningRigs[i].ip)
			if !miningRigs[i].Ping() {
				miningRigs[i].Restarter()
			}
		}
		log.Notice("Checking machines DONE")
	}

	work := func() {
		log.Notice("HELLO! I WILL KEEP YOUR MONEY MAKING MACHINES ONLINE!")
		// Run the check right away the first time
		check()
		timer := time.Duration(configFile.WaitSeconds) * time.Second
		log.Notice("Starting timer: ", timer)
		//Check the machines periodically
		gobot.Every(timer, check)
	}

	robot := gobot.NewRobot("RPiMinerHardReset", r, work)
	for _, rig := range miningRigs {
		robot.AddDevice(rig.pin)
	}

	robot.Start()
}
