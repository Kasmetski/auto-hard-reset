package main

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
)

//Ping the IP from Linux shell
func Ping(ip string) bool {
	out, _ := exec.Command("ping", ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		log.Println("HOST NOW FOUND - ", ip)
		return false
	}
	log.Println("Host ", ip, " is online")
	return true
}

//Restarter function logic
func Restarter(rig Rig) {
	log.Println("Restarting ", rig.name)
	rig.pin.Off()
	time.Sleep(5 * time.Second)
	rig.pin.On()
	time.Sleep(3 * time.Second)
	rig.pin.Off()
	time.Sleep(1 * time.Second)
	rig.pin.On()
	log.Println(rig.name, "restarted")
}

//Rig structure
type Rig struct {
	name string
	pin  *gpio.RelayDriver
	ip   string
	info string
}
