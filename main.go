package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	///MINING RIGS CONFIGURATION///
	var miningRigs [12]Rig //number of machines

	//EXAMPLE
	//miningRigs[num] = &Rig{"NAME", gpio.NewRelayDriver(r, "PIN-NUMBER OF GPIO"), "IP ADDRESS", "ADDITIONAL INFO"}
	miningRigs[0] = Rig{"machine 1", gpio.NewRelayDriver(r, "40"), "192.168.0.100", "R9 290's"}
	miningRigs[1] = Rig{"machine 2", gpio.NewRelayDriver(r, "38"), "192.168.0.101", "RX480's"}
	miningRigs[2] = Rig{"machine 3", gpio.NewRelayDriver(r, "37"), "192.168.0.102", "RX480's"}
	miningRigs[3] = Rig{"machine 4", gpio.NewRelayDriver(r, "36"), "192.168.0.103", "RX480's"}
	miningRigs[4] = Rig{"machine 5", gpio.NewRelayDriver(r, "35"), "192.168.0.104", "RX480's"}
	miningRigs[5] = Rig{"machine 6", gpio.NewRelayDriver(r, "33"), "192.168.0.105", "RX480's"}
	miningRigs[6] = Rig{"machine 7", gpio.NewRelayDriver(r, "32"), "192.168.0.106", "RX480's"}
	miningRigs[7] = Rig{"machine 8", gpio.NewRelayDriver(r, "31"), "192.168.0.107", "RX480's"}
	miningRigs[8] = Rig{"machine 9", gpio.NewRelayDriver(r, "29"), "192.168.0.108", "RX480's"}
	miningRigs[9] = Rig{"machine 10", gpio.NewRelayDriver(r, "22"), "192.168.0.109", "RX480's"}
	miningRigs[10] = Rig{"machine 11", gpio.NewRelayDriver(r, "18"), "192.168.0.110", "RX480's"}
	miningRigs[11] = Rig{"machine 12", gpio.NewRelayDriver(r, "16"), "192.168.0.111", "RX480's"}
	///END OF MINING RIG CONFIGURATION///

	LogMachines()

	work := func() {
		timer := 33 * time.Minute
		log.Notice("HELLO! I WILL KEEP YOUR MONEY MAKING MACHINES ONLINE!")
		log.Notice("Starting timer: ", timer)

		//Check the machines every 33 minutes
		gobot.Every(timer, func() {
			log.Notice("Checking machines: ")
			for i := 0; i < len(miningRigs); i++ {
				log.Notice("Ping miner: ", i, "name: ", miningRigs[i].name, "ip: ", miningRigs[i].ip)
				if !miningRigs[i].Ping() {
					miningRigs[i].Restarter()
				}
			}

			log.Notice("Checking machines DONE")
			log.Notice("Restarting timer")
		})
	}

	robot := gobot.NewRobot("RPiMinerHardReset",
		[]gobot.Connection{r},
		[]gobot.Device{miningRigs[0].pin},
		[]gobot.Device{miningRigs[1].pin},
		[]gobot.Device{miningRigs[2].pin},
		[]gobot.Device{miningRigs[3].pin},
		[]gobot.Device{miningRigs[4].pin},
		[]gobot.Device{miningRigs[5].pin},
		[]gobot.Device{miningRigs[6].pin},
		[]gobot.Device{miningRigs[7].pin},
		[]gobot.Device{miningRigs[8].pin},
		[]gobot.Device{miningRigs[9].pin},
		[]gobot.Device{miningRigs[10].pin},
		[]gobot.Device{miningRigs[11].pin},
		//IF YOU ADD MORE MACHINES ADD INFO HERE
		//[]gobot.Device{miningRigs[12].pin},
		work,
	)

	robot.Start()
}
