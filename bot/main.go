package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"math/rand"
	"time"

	"github.com/carrot/go-pinterest"
	libgiphy "github.com/sanzaru/go-giphy"

	forecast "github.com/mlbright/darksky/v2"

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

		//тут типа новый мембер
		if update.Message.NewChatMembers != nil {
			members := *update.Message.NewChatMembers
			for _, member := range members {
				// fmt.Println("вот тут лежит что-то ", member.UserName)
				reply := "Дарова! @" + member.UserName
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}

		}
		//мембер ушел
		if update.Message.LeftChatMember != nil {
			reply := "Скучать не буду! (да кого я обманываю) @" + update.Message.LeftChatMember.UserName
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}

		switch strings.ToLower(update.Message.Text) {
		case "слыш":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "слышу, слышу")
			bot.Send(msg)
		case "пикачу":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM3XisZmc0Vnfhgd2t5Ii8QuvolRxcAArkCAAI2diAO7XTrBWfSW4UYBA")
			bot.Send(msg)
		case "жопа марио":
			//идем на пинтерест
			client := pinterest.NewClient().RegisterAccessToken("AhWVuk8k2fUvGH1T_N3YBil8ysODFe4L7H5LHjpGfxGtP4Cu3ghDQDAAAj2mRn8VQyigwnsAAAAA")
			//берем айди и ищем
			pin, err := client.Pins.Fetch("748582769298226892")
			//ошибки в лог
			if err != nil {
				fmt.Println(err)
			}
			//вывод фоточки в чат
			msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, pin.Url)
			bot.Send(msg)
		case "пук":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANlXjAPO3YB8UrmMoDKSUwWstOGzvcAAjAAA8vnJA34Ty5P-pcKXBgE")
			bot.Send(msg)
		case "погода Краснодар":

			key := "c258e7938c7aca2458497faa4567b47d"
			lat := "45.0439"
			long := "38.9837"

			f, err := forecast.Get(key, lat, long, "now", forecast.SI, forecast.Russian)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %s\n", f.Timezone, f.Currently.Summary)
			fmt.Printf("Влажность: %.2f\n", f.Currently.Humidity)
			fmt.Printf("Температура: %.2f Celsius\n", f.Currently.Temperature)
			fmt.Printf("Скорость ветра: %.2f\n", f.Currently.WindSpeed)

			// text :=
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			// bot.Send(msg)
		case "жопа":
			//подключаемся к гифи
			giphy := libgiphy.NewGiphy("SgZU6aFi34lZB13cQ1Qv4lPqfxn3BuwH")
			//рандомим
			dataRandom, err := giphy.GetRandom("booty")
			if err != nil {
				fmt.Println("error:", err)
			}
			//урл нашей гифки
			url := dataRandom.Data.Image_original_url
			//запрос по урл
			gif, _ := http.Get(url)
			//загружаем файл
			file := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, tgbotapi.FileReader{
				Name:   url,
				Reader: gif.Body,
				Size:   gif.ContentLength,
			   })
			bot.Send(file)
		case "хорек":
			//подключаемся к гифи
			giphy := libgiphy.NewGiphy("SgZU6aFi34lZB13cQ1Qv4lPqfxn3BuwH")
			//рандомим
			dataRandom, err := giphy.GetRandom("ferret")
			if err != nil {
				fmt.Println("error:", err)
			}
			//урл нашей гифки
			url := dataRandom.Data.Image_original_url
			//запрос по урл
			gif, _ := http.Get(url)
			//загружаем файл
			file := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, tgbotapi.FileReader{
				Name:   url,
				Reader: gif.Body,
				Size:   gif.ContentLength,
			   })
			bot.Send(file)
		}
		//варианты внутри сообщения
		switch {
		case strings.Contains(strings.ToLower(update.Message.Text), "аниме"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM9XitUms2Xl_mcOU9cPtjOdbw0JDQAAvUDAALCGBUXIz-zor3JaQcYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "токсик"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANPXitXd_yFC9NwkrB_3b1aIUEcvh4AAgcEAALCGBUX9uiVkzQeVxsYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "бот понимает"):
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я понимаю, да")
			bot.Send(msg)
		}
		//варианты ответа на приветствие
		if strings.Contains(strings.ToLower(update.Message.Text), "привет") {
			rand.Seed(time.Now().UnixNano())
			answrs := []string{"Дарова","Куку епта","Мяу?"}
			num := rand.Intn(len(answrs))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answrs[num])
			bot.Send(msg)
			
		}
	}
}
