package main

import (
	"fmt"

	"github.com/gokyle/gopush/pushover"
)

//PushoverNotify using Pushover notification
func PushoverNotify(r *Rig) {
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
