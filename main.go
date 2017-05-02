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

	//store the info in miningRigs map
	//ToDo: config file
	miningRigs := make(map[string]Rig)

	miningRigs["rig1"] = Rig{"machine 1", gpio.NewRelayDriver(r, "38"), "192.168.0.100", "R9 290's"}
	miningRigs["rig2"] = Rig{"machine 2", gpio.NewRelayDriver(r, "40"), "192.168.0.111", "RX480's"}

	work := func() {
		log.Println("# Starting timer")

		//Check the machines every 10 minutes
		gobot.Every(10*time.Minute, func() {
			log.Println("# Checking machines: ")

			for key, value := range miningRigs {
				log.Println("# Ping miner: ", key, " ip: ", value.ip)
				if !Ping(value.ip) {
					log.Println("### HOST NOT FOUND - ", value.name)
					Restarter(miningRigs[key])
				} else {
					log.Println("# HOST IS ONLINE - ", value.name)
				}
			}
			log.Println("# Checking machines DONE")
			log.Println("# Starting timer")
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
