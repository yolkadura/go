package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

//точка входа
func main() {
	botToken := "753203683:AAF7UUwaVSihMMa_R2qci-QZqkRgQ3tV1HE"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken

	for ;; { //бесконечный цикл обновлений
		updates, err := getUpdates(botUrl)
			if err != nil {
				log.Println("something wrong: ", err.Error())
			}
		fmt.Println(updates)
	}
}

//запросы обновлений
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
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
func respond() {

}