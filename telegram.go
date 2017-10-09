package main

import "time"
import "github.com/go-telegram-bot-api/telegram-bot-api"
import "strconv"
import "fmt"

//TelegramBot is the main function used to communicate with Telegram chat and remote control
func TelegramBot(rigs []Rig) {
	bot, err := tgbotapi.NewBotAPI(Config.TgAPIKey)
	if err != nil {
		log.Fatal("Problem with the API key. Fix the problem or disable the TelegramBot from the config file")
	}

	bot.Debug = false

	log.Noticef("TelegramBot Authorteleized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Noticef("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			commandArgs := update.Message.CommandArguments()

			if update.Message.From.UserName != Config.TgAdminUserName {
				msg.Text = "You are not my master"
			} else {
				switch update.Message.Command() {
				case "help":
					msg.Text = "there is no help for the people here\nbut you can try /status /config /turnon /turnoff /restart /ping"

				case "status":
					msg.Text = "I'm fine. Thanks for asking"

				case "config":
					msg.Text = handleConfig()

				case "ping":
					msg.Text = handlePing(rigs, commandArgs)

				case "restart":
					msg.Text = handleRestart(rigs, commandArgs)
				case "turnon":
					msg.Text = handleTurnOn(rigs, commandArgs)

				case "turnoff":
					msg.Text = handleTurnOff(rigs, commandArgs)

				default:
					msg.Text = "I don't know that command, try /help"
				}
			}

			bot.Send(msg)
		}
	}
}

//handleConfig return string with raw data from config file
func handleConfig() string {
	return fmt.Sprintf("Config:\nLogging: %t\nStartupcheck: %t\nTimer: %d\nMiners:\n%+v",
		Config.Log, Config.StartupCheck, Config.WaitSeconds, Config.Miners)
}

//handlePing return information about the ping response from machine
func handlePing(rigs []Rig, commandArgs string) (s string) {
	if commandArgs != "" {
		args, _ := strconv.ParseInt(commandArgs, 10, 8)
		s = fmt.Sprintf("Machine online status: %t", rigs[args].Ping())
	} else {
		s = "Provide arguments"
	}

	return
}

//handleRestart restarts the machine and returns string
func handleRestart(rigs []Rig, commandArgs string) (s string) {
	if commandArgs != "" {
		args, _ := strconv.ParseInt(commandArgs, 10, 8)
		rigs[args].Restarter()
		s = fmt.Sprintf("Machine %d was restarted", args)
	} else {
		s = "Provide arguments"
	}

	return
}

//handleTurnOn turns on the machine and returns string
func handleTurnOn(rigs []Rig, commandArgs string) (s string) {
	if commandArgs != "" {
		args, _ := strconv.ParseInt(commandArgs, 10, 8)
		rigs[args].TurnOn()
		s = fmt.Sprintf("Machine %d was turnedoff", args)
	} else {
		s = "Provide arguments"
	}

	return
}

//handleTurnOff restarts the machine and returns string
func handleTurnOff(rigs []Rig, commandArgs string) (s string) {
	if commandArgs != "" {
		args, _ := strconv.ParseInt(commandArgs, 10, 8)
		rigs[args].ForceShutDown()
		s = fmt.Sprintf("Machine %d was turnedoff", args)
	} else {
		s = "Provide arguments"
	}

	return
}
