package types

import (
	"encoding/json"
	"net/url"
)

type Update struct {
	UpdateID      int64          `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	ChannelPost   *ChannelPost   `json:"channel_post,omitempty"`
}

type Message struct {
	MessageID int64      `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int64      `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}

type ChannelPost struct {
	MessageID int64  `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type SendMessage struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	MessageID int64  `json:"message_id"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
}

type From struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID                          int64  `json:"id"`
	FirstName                   string `json:"first_name"`
	UserName                    string `json:"username"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

type Entities struct {
	Type   string   `json:"type"`
	Offset int64    `json:"offset"`
	Length int64    `json:"length"`
	Url    *url.URL `json:"url,omitempty"`
}

type SendMessageRequest struct {
	ChatID      int64           `json:"chat_id" structs:"chat_id"`
	Text        string          `json:"text" structs:"text"`
	ParseMode   string          `json:"parse_mode,omitempty" structs:"parse_mode,omitempty"`
	ReplyMarkup json.RawMessage `json:"reply_markup,omitempty" structs:"reply_markup,omitempty"`
	Entities    []Entities      `json:"entities,omitempty" structs:"entities,omitempty"`
}

type SendPhotoRequest struct {
	ChatID      int64           `json:"chat_id" structs:"chat_id"`
	Photo       string          `json:"photo" structs:"photo"`
	Caption     string          `json:"caption,omitempty" structs:"caption,omitempty"`
	ParseMode   string          `json:"parse_mode,omitempty" structs:"parse_mode,omitempty"`
	ReplyMarkup json.RawMessage `json:"reply_markup,omitempty" structs:"reply_markup,omitempty"`
	Entities    []Entities      `json:"entities,omitempty" structs:"entities,omitempty"`
}

type TelegramError struct {
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
	Ok          bool   `json:"ok"`
}

type CallbackQuery struct {
	ID           string  `json:"id"`
	From         From    `json:"from"`
	Message      Message `json:"message"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
}
