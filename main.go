package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	//Read configuration file
	Config = ReadConfig()

	//parse machines to []Rig struct
	miningRigs := make([]Rig, 0)
	for _, m := range Config.Miners {
		log.Notice("minerConfig:", m)
		miningRigs = append(miningRigs, Rig{m.Name, gpio.NewRelayDriver(r, m.Pin), m.IP, m.Info})
	}

	log.Notice("Configured rigs: ", len(miningRigs))

	//Logging machines in two outputs - console & external file
	if Config.Log {
		go LogMachines()
	}

	if Config.TgBotActivate {
		go TelegramBot(miningRigs)
	}
	//Gobot work func
	work := func() {
		log.Notice("HELLO! I WILL KEEP YOUR MINING RIGS ONLINE!")

		//Check machines on startup without waiting the timer. Use with caution. After a power failure, RPI could be ready faster than your machines and start restarting them without need.
		if Config.StartupCheck {
			CheckMachines(miningRigs)
		}

		timer := time.Duration(Config.WaitSeconds) * time.Second
		log.Notice("Starting timer: ", timer)

		//Check the machines periodically
		gobot.Every(timer, func() {
			CheckMachines(miningRigs)
		})
	}

	robot := gobot.NewRobot("auto-hard-reset", r, work)
	for _, rig := range miningRigs {
		robot.AddDevice(rig.pin)
	}

	robot.Start()
}
