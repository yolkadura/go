package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	//это была строка без сокса bot, err := tgbotapi.NewBotAPI("753203683:AAF7UUwaVSihMMa_R2qci-QZqkRgQ3tV1HE")
	bot, err := tgbotapi.NewBotAPIWithClient("753203683:AAF7UUwaVSihMMa_R2qci-QZqkRgQ3tV1HE", httpClient)
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

		if update.Message.Text == "слыш" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "слышу, слышу")
			bot.Send(msg)

		}

		//тут типа новый мембер, но хз работает ли
		var reply string

		if update.Message.NewChatMembers != nil {
			// В чат вошел новый пользователь
			// Поприветствуем его
			reply = fmt.Sprintf("Даров @%s! Семки есть?", update.Message.NewChatMembers)
		}
		if reply != "" {
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			// и отправляем его
			bot.Send(msg)
		}

	}
}
