package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//точка входа
func main() {
	botToken := "753203683:AAF7UUwaVSihMMa_R2qci-QZqkRgQ3tV1HE"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0

	for { //бесконечный цикл обновлений
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("something wrong: ", err.Error())
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

//запросы обновлений
func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //отложеное закрытие тела
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse //массив результатов
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

//ответ на обновления
func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
