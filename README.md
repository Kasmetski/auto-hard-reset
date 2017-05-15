# RPi Miner Hard-reset
Simple Go program for auto-hard-reset computer with GPIO
#### Still in progress
![console](screenshot.png)

First commit is the prototype. I'm using 5V relay and checking the miners with ping command. If there is no answer - hard-reset with the cables.

### Logic
 * Ping miners every 10 minutes
 * if offline > send signal for 5 sec, pause 5 sec and send signal again for 0.108 second

### How-to

There is still no config files (searching a good solution for Go). Machines are stored in maps `miningRigs[key]` with `Rig` struct looking like this

```
type Rig struct {
	name string
	pin  *gpio.RelayDriver
	ip   string
	info string
}
```

To add new machines open `main.go`, edit the examples and add more if you need.
```
miningRigs["rig1"] = &Rig{"machine 1", gpio.NewRelayDriver(r, "38"), "192.168.0.100", "R9 290's"}
```
```
miningRigs["KEY-NAME"] = &Rig{"NAME", gpio.NewRelayDriver(r, "PIN-NUMBER OF RASPBERRY PY"), "LOCAL IP ADDRESS", "ADDITIONAL INFO"}
```

### ToDo
* Config file for all miners
* JSON check/
* instructions
* statistics