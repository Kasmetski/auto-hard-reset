# RPi miner auto-hard-reset
Simple Go program for auto-hard-reset computer with Raspberry Pi or other mini-computer
#### Still in progress
![console](screenshot.png)

First commit is the prototype. I'm using 5V relay and checking the miners with ping command. If there is no answer - hard-reset with the GPIO.

### Requirements
* Raspberry Pi
* Golang >= 1.8.0

### Logic
 * Ping miners every 33 minutes
 * if offline > send signal for 5 sec (turn off pc), pause 5 sec(wait) and send signal again for 0.108 second (turn on pc)

### How-to
I'm using Raspberry Pi with 5V relay. Soon I will add detailed instruction but this is the basics.
![console](raspberrypi-5v-relay.jpeg)
Soon simple configuration and binary files.

Use  `go get -u -v github.com/kasmetski/auto-hard-reset` instead of `git clone`

Machines are stored in an array `miningRigs[]`

To add new machines open `main.go`, edit the examples and delete the unnecessary ones.
First write how many machines you will control
```
	var miningRigs [12]Rig //number of machines
```
After that add/edit the machines like that
```
miningRigs[0] = Rig{"machine 1", gpio.NewRelayDriver(r, "38"), "192.168.0.100", "R9 290's"}
```

```
miningRigs[num] = Rig{"NAME", gpio.NewRelayDriver(r, "PIN-NUMBER OF RASPBERRY PI"), "LOCAL IP ADDRESS", "ADDITIONAL INFO"}
```

and in the end add the new machines or delete unnecessary in the last function

```
robot := gobot.NewRobot("RPiMinerHardReset",
		[]gobot.Connection{r},
		[]gobot.Device{miningRigs[0].pin},
		[]gobot.Device{miningRigs[1].pin},
		[]gobot.Device{miningRigs[2].pin},
		[]gobot.Device{miningRigs[3].pin},
		[]gobot.Device{miningRigs[4].pin},
        .....
        and so on
```

### Build
If you are building on your Raspberry Pi, type `go build *.go` in the folder.
If you are building on your workstation type `GOARM=6(or 7) GOARCH=arm GOOS=linux go build *.go`
##### GOARM=6 (Raspberry Pi A, A+, B, B+, Zero) GOARM=7 (Raspberry Pi 2, 3)

### ToDo
* conf.file
* web interface
* JSON-check
* instructions
* statistics