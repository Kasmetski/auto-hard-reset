package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	miningRigs := make(map[string]*Rig)

	///MINING RIGS CONFIGURATION///
	//EXAMPLE - miningRigs["KEY-NAME"] = &Rig{"NAME", gpio.NewRelayDriver(r, "PIN-NUMBER OF RASPBERRY PY"), "LOCAL IP ADDRESS", "ADDITIONAL INFO"}
	miningRigs["rig1"] = &Rig{"machine 1", gpio.NewRelayDriver(r, "38"), "192.168.0.100", "R9 290's"}
	miningRigs["rig2"] = &Rig{"machine 2", gpio.NewRelayDriver(r, "40"), "192.168.0.101", "RX480's"}
	///END OF MINING RIG CONFIGURATION///

	work := func() {
		log.Println("# Starting timer")
		//Check the machines every 10 minutes
		gobot.Every(10*time.Minute, func() {
			//logging
			t := time.Now().String()
			fname := fmt.Sprintf("./auto-hard-reset-log-%s.txt", t[:10])

			file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
			if err != nil {
				return
			}
			defer file.Close()
			log.SetOutput(file)

			//checking machines
			log.Println("# Checking machines: ")

			for k, v := range miningRigs {
				fmt.Println("# Ping miner: ", k, " ip: ", v.ip)

				if !v.Ping() {
					log.Println("##### HOST NOT FOUND - ", v.name)
					miningRigs[k].Restarter()
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
		//IF YOU ADD MORE MACHINES ADD INFO HERE
		//[]gobot.Device{miningRigs["rig3"].pin},
		work,
	)

	robot.Start()
}
