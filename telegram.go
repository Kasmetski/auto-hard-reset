package main

import "time"
import "github.com/go-telegram-bot-api/telegram-bot-api"
import "strconv"
import "fmt"

func TelegramBot(rigs []Rig) {
	machines := rigs
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
				args, err := strconv.ParseInt(commandArgs, 10, 32)
				if err != nil {
					log.Error(err)
				}

				switch update.Message.Command() {
				case "help":
					msg.Text = "there is no help for the people here\nbut you can try /status /miners /turnon /turnoff /restart /ping"

				case "status":
					msg.Text = "I'm fine. Thanks for asking"

				case "miners":
					msg.Text = printConf()

				case "ping":
					msg.Text = fmt.Sprintf("Machine online status: %t", machines[args].Ping())

				case "restart":
					machines[args].Restarter()
					msg.Text = fmt.Sprintf("Machine %d was restarted", args)

				case "turnon":
					machines[args].TurnOn()
					msg.Text = fmt.Sprintf("Machine %d started", args)

				case "turnoff":
					machines[args].ForceShutDown()
					msg.Text = fmt.Sprintf("Machine %d was turnedoff", args)

				default:
					msg.Text = "I don't know that command, try /help"
				}
			}

			bot.Send(msg)
		}
	}
}

//return string with raw data from config file
func printConf() string {
	return fmt.Sprintf("Config:\n%+v", Config)
}
