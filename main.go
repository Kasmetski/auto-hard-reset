package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var miningRigs [5]string

func main() {
	miningRigs[0] = "192.168.0.111"

	r := raspi.NewAdaptor()
	//led := gpio.NewLedDriver(r, "7")
	relay := gpio.NewRelayDriver(r, "7")
	work := func() {
		gobot.Every(10*time.Minute, func() {
			fmt.Println("LED TOGGGLEEE")
			if !Ping(miningRigs[0]) {
				relay.Off()
				time.Sleep(5 * time.Second)
				relay.On()
				time.Sleep(3 * time.Second)
				relay.Off()
				time.Sleep(1 * time.Second)
				relay.On()
			}
		})
	}

	// result := Ping(miningRigs[1])
	// fmt.Println(result)
	robot := gobot.NewRobot("RPiMinerHardReset",
		[]gobot.Connection{r},
		[]gobot.Device{relay},
		work,
	)

	robot.Start()
}

//Restarter -
