package tgtypes

type Update struct {
	Message  Message `json:"message"`
	UpdateID int     `json:"update_id"`
}

type Message struct {
	Chat struct {
		FirstName string `json:"first_name"`
		ID        int    `json:"id"`
		Type      string `json:"type"`
		Username  string `json:"username"`
	} `json:"chat"`
	Date     int `json:"date"`
	Entities []struct {
		Length int    `json:"length"`
		Offset int    `json:"offset"`
		Type   string `json:"type"`
	} `json:"entities"`
	From struct {
		FirstName    string `json:"first_name"`
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		LanguageCode string `json:"language_code"`
		Username     string `json:"username"`
	} `json:"from"`
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
}

type OutgoingMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
