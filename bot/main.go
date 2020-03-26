package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	libgiphy "github.com/sanzaru/go-giphy"

	forecast "github.com/mlbright/darksky/v2"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"

	// geo "github.com/codingsince1985/geo-golang/yandex"
)

var (	
giphy *libgiphy.Giphy //объявляем апи для гифи
bot *tgbotapi.BotAPI //объявляем апи для телеги
// f *forecast.
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

	//чтение файла
	file, err := ioutil.ReadFile("/home/yolka/go/woodman_api.txt")
	if err != nil {
		//return
		fmt.Println("Error", err)
	}
	//перевод файла в строку
	yourawesomekey := string(file)

	//подключение к апи телеги
	//это была строка без сокса bot, err := tgbotapi.NewBotAPI(yourawesomekey)
	bot, err = tgbotapi.NewBotAPIWithClient(yourawesomekey, httpClient)
	if err != nil {
		log.Panic(err)
	}

	//подключаемся к апи гифи
	giphy = libgiphy.NewGiphy("SgZU6aFi34lZB13cQ1Qv4lPqfxn3BuwH")

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
				reply := "Дарова! @" + member.UserName
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
		case "пикачу":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM3XisZmc0Vnfhgd2t5Ii8QuvolRxcAArkCAAI2diAO7XTrBWfSW4UYBA")
			bot.Send(msg)
		case "пук":
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANlXjAPO3YB8UrmMoDKSUwWstOGzvcAAjAAA8vnJA34Ty5P-pcKXBgE")
			bot.Send(msg)
		case "погода крас":

			pogoda("45.0439","38.9837",update.Message.Chat.ID)
		case "погода питер":
			pogoda("59.8211","30.3882",update.Message.Chat.ID)
		case "погода барнаул":
			pogoda("53.3509","83.7563",update.Message.Chat.ID)
		}

		//триггеры на слова
		switch {
		case strings.Contains(strings.ToLower(update.Message.Text), "привет"): //варианты ответа на приветствие
			rand.Seed(time.Now().UnixNano())
			answrs := []string{"Дарова", "Куку епта", "Мяу?"}
			num := rand.Intn(len(answrs))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answrs[num])
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "аниме"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAM9XitUms2Xl_mcOU9cPtjOdbw0JDQAAvUDAALCGBUXIz-zor3JaQcYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "токсик"):
			msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAANPXitXd_yFC9NwkrB_3b1aIUEcvh4AAgcEAALCGBUX9uiVkzQeVxsYBA")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "чудовище"):
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "А на себя давно в зеркало смотрел?")
			bot.Send(msg)
		case strings.Contains(strings.ToLower(update.Message.Text), "покемон"):
			findgif("pokemon",update.Message.Chat.ID)
		case strings.Contains(strings.ToLower(update.Message.Text), "жопа"):
			findgif("booty",update.Message.Chat.ID)
		case strings.Contains(strings.ToLower(update.Message.Text), "хорек"):
			findgif("ferret",update.Message.Chat.ID)
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
func pogoda(lat string, long string, chatid int64) {
	key := "c258e7938c7aca2458497faa4567b47d"
	f, err := forecast.Get(key, lat, long, "now", forecast.SI, forecast.Russian)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(f)
	text := fmt.Sprintf("Часовой пояс: %s\nПогода: %s\nВлажность: %.0f %%\nТемпература: %.2f C\nСкорость ветра: %.2f м/с\n", f.Timezone, f.Currently.Summary, f.Currently.Humidity*100, f.Currently.Temperature, f.Currently.WindSpeed)
	msg := tgbotapi.NewMessage(chatid, text)
	bot.Send(msg)
}