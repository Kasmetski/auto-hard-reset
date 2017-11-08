package main

import (
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

//CheckMachines - ping machines, if there is no responce >> hard-reset
func CheckMachines(r []Rig) {
	log.Notice("Checking machines: ")

	for i := 0; i < len(r); i++ {
		log.Notice("Ping machine: ", r[i].name, "ip: ", r[i].ip)
		if !r[i].Ping() {
			r[i].Restarter()
		}
	}

	log.Notice("Checking machines DONE\n----------------------")
	log.Notice("Starting timer")
}

//Ping IP from Linux shell
func (r *Rig) Ping() bool {
	out, _ := exec.Command("ping", r.ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "100% packet loss") {
		log.Error("HOST NOT FOUND: ", r.name, r.ip)
		return false
	}

	log.Notice("HOST IS ONLINE: ", r.name)
	return true
}

//ForceShutDown machine
func (r *Rig) ForceShutDown() {
	r.pin.Off()
	time.Sleep(5 * time.Second)
	r.pin.On()
}

//TurnOn machine
func (r *Rig) TurnOn() {
	r.pin.Off()
	time.Sleep(108 * time.Millisecond)
	r.pin.On()
}

//Restarter function logic
func (r *Rig) Restarter() {
	log.Warning("Restarting: ", r.name)

	Notify(r)

	r.ForceShutDown()
	time.Sleep(5 * time.Second)
	r.TurnOn()

	log.Warning("Machine restarted: ", r.name)
}
