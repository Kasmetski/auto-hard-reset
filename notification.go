package main

import (
	"fmt"

	"bitbucket.org/kisom/gopush/pushover"
)

//Notify using Pushover notification
func Notify(r *Rig) {
	if Config.Pushover == true {
		identity := pushover.Authenticate(
			Config.PushoverToken,
			Config.PushoverUser,
		)

		message := fmt.Sprint("Force rebooting ", r.name)
		sent := pushover.Notify(identity, message)
		if !sent {
			log.Error("[!]Pushover notification failed.")
		} else {
			log.Notice("Pusherover notification sent")
		}
	}
}
