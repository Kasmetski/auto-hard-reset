package main

import (
	"fmt"
	"os"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("auto-hard-reset-log")

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
