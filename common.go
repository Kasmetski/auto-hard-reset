package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	logging "github.com/op/go-logging"

	"gobot.io/x/gobot/drivers/gpio"
)

//Rig structure
type Rig struct {
	name string
	pin  *gpio.RelayDriver
	ip   string
	info string
}

//Ping IP from Linux shell
func (r *Rig) Ping() bool {
	out, _ := exec.Command("ping", r.ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Host Unreachable") {
		log.Error("HOST NOT FOUND: ", r.name, r.ip)
		return false
	}
	log.Notice("HOST IS MAKING MONEY: ", r.name)
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
	r.ForceShutDown()
	time.Sleep(5 * time.Second)
	r.TurnOn()
	log.Warning("Machine restarted: ", r.name)

}

//LogMachines - function for basic logging
func LogMachines() {
	t := time.Now().Format("2006-01-02-15-04-05")
	fname := fmt.Sprintf("./auto-hard-reset-log-%s.txt", t)

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return
	}
	//defer f.Close()

	b1format := logging.MustStringFormatter(`%{time:2006-01-02-15:04:05.000} ▶ %{level:.5s} %{id:03x} ▶ %{message}`)
	b2format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} ▶ %{level:.5s} %{id:03x} ▶ %{color:reset} %{message}`)

	b1 := logging.NewLogBackend(f, "", 0)
	b2 := logging.NewLogBackend(os.Stderr, "", 0)

	//Formating messages
	b1Formater := logging.NewBackendFormatter(b1, b1format)
	b2Formater := logging.NewBackendFormatter(b2, b2format)

	//Sending errors to b1logger
	b1Leveled := logging.AddModuleLevel(b1Formater)
	b1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(b1Leveled, b2Formater)
}
