package main

import (
	"fmt"
	"log"
	"os"
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
func (r *Rig) Ping() bool {
	out, _ := exec.Command("ping", r.ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Host Unreachable") {
		return false
	}
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
	log.Println("### Restarting ", r.name)
	r.ForceShutDown()
	time.Sleep(5 * time.Second)
	r.TurnOn()
}

//LogMachines -
func LogMachines() {
	t := time.Now().String()
	fname := fmt.Sprintf("./auto-hard-reset-log-%s.txt", t[:10])

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return
	}
	//defer file.Close()
	log.SetOutput(file)
}
