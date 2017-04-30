package main

import (
	"os/exec"
	"strings"
)

//Ping an IP
func Ping(ip string) bool {
	out, _ := exec.Command("ping", ip, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		return false
	}

	return true

}
