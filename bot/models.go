package main


type Update struct {
	Update_id int			`json:"update_id"`
	Message Message			`json:"message"`
}

type Message struct {
	Chat Chat				`json:"chat"`
	Text string				`json:"text"`
}

type Chat struct {
	ChatId int				`json:"chat"`
}

type RestResponse struct {
	Result []Update			`json:"result"`
}