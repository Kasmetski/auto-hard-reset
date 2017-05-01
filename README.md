# RPi Miner Hard-reset
Simple Go program for hard-reset computer with GPIO
#### Still in progress
First commit is the prototype. I'm using 5V relay and checking the miners with ping command. If there is no answer - hard-reset with the cables.

### Logic
 * Ping miners every 10 minutes
 * if offline > send signal for 5 sec, pause 3 sec and send signal again for 1 second

### ToDo
* Config file for all miners
* JSON check/
* log file
* instructions