package main

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	//store the info in miningRigs map | using empty ip for testing purpose
	//ToDo: config file
	miningRigs := make(map[string]Rig)
	miningRigs["rig1"] = Rig{"machine 1", gpio.NewRelayDriver(r, "38"), "192.168.0.111", "6 x RX480"}
	miningRigs["rig2"] = Rig{"machine 2", gpio.NewRelayDriver(r, "40"), "192.168.0.111", "6 x RX470"}

	work := func() {
		log.Println("Starting timer")
		//Check the machines every 10 minutes
		gobot.Every(10*time.Minute, func() {
			log.Println("Checking machines: ")
			for key, value := range miningRigs {
				log.Println("Ping miner: ", key, " ip: ", value.ip)
				if !Ping(value.ip) {
					Restarter(miningRigs[key])
				}
			}

		})
	}
	robot := gobot.NewRobot("RPiMinerHardReset",
		[]gobot.Connection{r},
		[]gobot.Device{miningRigs["rig1"].pin},
		[]gobot.Device{miningRigs["rig2"].pin},
		work,
	)

	robot.Start()
}
