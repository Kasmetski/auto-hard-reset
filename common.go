package main

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
)

//Rig structure
type Rig struct {
	name string
	pin  *gpio.RelayDriver
	ip   string
	info string
}

//Ping the IP from Linux shell
func Ping(ip string) bool {
	out, _ := exec.Command("ping", ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		return false
	}
	return true
}

//ForceShutDown machine
func ForceShutDown(r Rig) {
	r.pin.Off()
	time.Sleep(5 * time.Second)
	r.pin.On()
}

//TurnOn machine
func TurnOn(r Rig) {
	r.pin.Off()
	time.Sleep(1 * time.Second)
	r.pin.On()
}

//Restarter function logic
func Restarter(r Rig) {
	log.Println("### Restarting ", r.name)
	ForceShutDown(r)
	time.Sleep(3 * time.Second)
	TurnOn(r)
}
