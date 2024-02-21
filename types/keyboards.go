package types

import "encoding/json"

type ReplyMarkup struct {
	InlineKeyboardMarkup [][]InlineKeyboardButton `json:"inline_keyboard" structs:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text" structs:"text"`
	URL          string `json:"url,omitempty" structs:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty" structs:"callback_data,omitempty"`
}

func (i *ReplyMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(*i)
}
