package main

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"os"
)

var channelSecret string
var channelToken string
var bot *messaging_api.MessagingApiAPI

func StartBot() {
	channelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	channelToken = os.Getenv("LINE_CHANNEL_TOKEN")

	var err error
	bot, err = messaging_api.NewMessagingApiAPI(channelToken)
	if err != nil {
		logger.Fatal(err)
	}
}
