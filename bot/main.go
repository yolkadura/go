package main

import {
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
}

//точка входа
func main() {
	botToken := "753203683:AAFDp7VhdHWWblFAQeZf1u8qKourGPaQyJY"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken

	for ;; { //бесконечный цикл обновлений
		updates, err := getUpdates(botUrl)
			if err != nil {
				return log.Println("someting wrong: ", err.Error())
			}
		fmt.Println(updates)
	}
}

//запросы обновлений
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
		if err != nil {
			return nill, err
		}
	defer resp.Body.Close() //отложеное закрытие тела
	body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nill, err
		}
	var restResponse RestResponse //массив результатов
	err := json.Unmarshal(body, &restResponse)
		if err != nil {
			return nill, err
		}
	return restResponse.Result, nil
}

//ответ на обновления
func respond() {

}