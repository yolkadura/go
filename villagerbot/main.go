package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os" //для прокси
	"strings"
	"time"

	libgiphy "github.com/sanzaru/go-giphy"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy" //для прокси
)

var (
	giphy *libgiphy.Giphy  //объявляем апи для гифи
	bot   *tgbotapi.BotAPI //объявляем апи для телеги
)

const PROXY_ADDR = "213.183.59.195:1080"

func main() {

	//логпассы сокса - прокся
	auth := proxy.Auth{
		User:     "banya",
		Password: "supersiski",
	}

	//настройка подключения - прокся
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, &auth, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	// настройки клиента - прокся
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial
	httpClient := &http.Client{Transport: httpTransport}

	yourawesomekey := readToken("/home/yolka/keys/villager_api.txt")

	//подключение к апи телеги
	//это была строка без сокса
	// bot, err := tgbotapi.NewBotAPI(yourawesomekey)
	//подключение через проксю
	bot, err = tgbotapi.NewBotAPIWithClient(yourawesomekey, httpClient)
	if err != nil {
		log.Panic(err)
	}

	giphykey := readToken("/home/yolka/keys/giphy_api.txt")

	//подключаемся к апи гифи
	giphy = libgiphy.NewGiphy(giphykey)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	//бесконечный цикл получения сообщений
	for update := range updates {
		if update.Message == nil { // игнорируем пустые сообщения
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//эхо
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID
		// bot.Send(msg)

		//тут типа новый мембер
		if update.Message.NewChatMembers != nil {
			members := *update.Message.NewChatMembers
			for _, member := range members {
				// fmt.Println("вот тут лежит что-то ", member.UserName)
				reply := "Добро пожаловать, сосед @" + member.UserName + "!\nУ нас тут самый ламповый чатик по кроссингу, поэтому веди себя хорошо!\n/fc - первым делом можешь заполнить форму дружбы или добавить друзей из таблицы кодов!\nТак же по команде /help могу подсказать разного по основам игры!\nПиратство осуждаем D:<\nПОТРУБИМ :D"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}
		}
		//мембер ушел
		if update.Message.LeftChatMember != nil {
			reply := "чао пака"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
		//команды
		switch strings.ToLower(update.Message.Text) {
		case "/help":
			reply := "/gif - гифка по АС\n/fc - форма и таблица кодов дружбы\nПиксельарты и шмоточки тут - @acpixelart\nГайды по игре тут - @BiblioAC"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "/gif":
			findgif("animalcrossing", update.Message.Chat.ID)
		case "/fc", "фк":
			reply := "Вот тут у нас форма с кодами - https://forms.gle/rmfhkMiCKFMgCEii6\nА вот тут таблица - https://docs.google.com/spreadsheets/d/1LfrnZJmgwGPYxJaBLtRnjZMNDnevyjXmC9ZDjqBS2Bo/edit?usp=sharing"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "/art", "арт":
			reply := "Пиксельарты и дизайн шмоточек по кроссенгу тут - @acpixelart"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "hui":
			msg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(msg)
		}

		//триггеры на слова
		switch {
		case strings.Contains(strings.ToLower(update.Message.Text), "привет"): //варианты ответа на приветствие
			rand.Seed(time.Now().UnixNano())
			answrs := []string{"CAACAgQAAxkBAAMKXnxGj_BlQmHNrq_wlFqd0J68rIkAAnwAA0csHgABCpRsAcnfZ0oYBA", "CAACAgQAAxkBAAMQXnxH7FM6_BXei2PH18wf9FDD_GYAApwAA0csHgABPfqm8NXFjo4YBA"}
			num := rand.Intn(len(answrs))
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, answrs[num])
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "аниме"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM9XitUms2Xl_mcOU9cPtjOdbw0JDQAAvUDAALCGBUXIz-zor3JaQcYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "токсик"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANPXitXd_yFC9NwkrB_3b1aIUEcvh4AAgcEAALCGBUX9uiVkzQeVxsYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "хуй"):
			msg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(msg)
		}
	}
}

func findgif(tag string, chatid int64) {
	//рандомим
	dataRandom, err := giphy.GetRandom(tag)
	if err != nil {
		fmt.Println("error:", err)
	}
	//урл нашей гифки
	url := dataRandom.Data.Image_original_url
	//запрос по урл
	gif, _ := http.Get(url)
	//загружаем файл
	file := tgbotapi.NewDocumentUpload(chatid, tgbotapi.FileReader{
		Name:   url,
		Reader: gif.Body,
		Size:   gif.ContentLength,
	})
	bot.Send(file)
}
func readToken(path string) string {
	//чтение файла
	file, err := ioutil.ReadFile(path)
	if err != nil {
		//return
		fmt.Println("Error", err)
	}
	//перевод файла в строку
	yourawesomekey := string(file)
	return yourawesomekey
}
