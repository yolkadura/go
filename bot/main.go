package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
)

const PROXY_ADDR = "213.183.59.195:1080"

func main() {

	//логпассы сокса
	auth := proxy.Auth{
		User:     "banya",
		Password: "supersiski",
	}

	//настройка подключения
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, &auth, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	// настройки клиента
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial
	httpClient := &http.Client{Transport: httpTransport}

	//это была строка без сокса bot, err := tgbotapi.NewBotAPI("")
	bot, err := tgbotapi.NewBotAPIWithClient("", httpClient)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // игнорируем пустые сообщения
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//эхо
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID
		// bot.Send(msg)

		//тут типа новый мембер, но хз работает ли
		if update.Message.NewChatMembers != nil {
			// reply := "Дарова! " + update.Message.
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			// bot.Send(msg)
		}

		if update.Message.LeftChatMember != nil {
			reply := "Ну и иди лесом!" + update.Message.From.UserName
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}

		switch update.Message.Text {
		case "слыш":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "слышу, слышу")
			bot.Send(msg)
		case "жопка":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "сам такой")
			bot.Send(msg)
		case "пикачу":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM3XisZmc0Vnfhgd2t5Ii8QuvolRxcAArkCAAI2diAO7XTrBWfSW4UYBA")
			bot.Send(msg)
		case "токсик":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANPXitXd_yFC9NwkrB_3b1aIUEcvh4AAgcEAALCGBUX9uiVkzQeVxsYBA")
			bot.Send(msg)
		case "жопа":
			msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, "https://66.media.tumblr.com/c8878079606e35ea63e7bf10322f3200/tumblr_phtfwhni8Z1te51fyo1_540.png")
			bot.Send(msg)
		}

		if strings.Contains(update.Message.Text, "аниме") == true {
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM9XitUms2Xl_mcOU9cPtjOdbw0JDQAAvUDAALCGBUXIz-zor3JaQcYBA")
			bot.Send(msg)
		}

	}
}
