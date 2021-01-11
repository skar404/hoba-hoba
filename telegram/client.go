package telegram

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/skar404/hoba-hoba/requests"
)

type MessageReq struct {
	ChatId           int    `json:"chat_id"`
	Mode             string `json:"parse_mode"`
	Text             string `json:"text,omitempty"`
	ReplyToMessageId *int   `json:"reply_to_message_id,omitempty"`

	//Caption string `json:"caption,omitempty"`
	//Photo   string `json:"photo,omitempty"`
	//Video   string `json:"video,omitempty"`
}

var Client = requests.RequestClient{
	// Указан url локального Telegram Server API чтобы убрать ограничения с размером файла для BOT API
	Url:     fmt.Sprintf("https://telegram-api.y.ulock.org/bot%s/", os.Getenv("TG_TOKEN")),
	Timeout: 30 * time.Second,
	Header: map[string][]string{
		"Content-Type": {"application/json"},
		"charset":      {"utf-8"},
	},
}

func SendAudio(chatId int, fileName string, file []byte) (int, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	err := writer.WriteField("chat_id", strconv.Itoa(chatId))
	if err != nil {
		return 0, err
	}

	part, err := writer.CreateFormFile("audio", fileName)
	if err != nil {
		return 0, err
	}

	_, err = part.Write(file)
	if err != nil {
		return 0, err
	}

	err = writer.Close()
	if err != nil {
		return 0, err
	}

	req := requests.Request{
		Method: http.MethodPost,
		Uri:    "sendAudio",
		Body:   &body,
		Header: map[string][]string{
			"Content-Type": {writer.FormDataContentType()},
		},
		Flags: requests.RequestFlags{
			IsBodyString: true,
		},
	}

	resData := SendAudioRes{}
	res := requests.Response{
		Struct: &resData,
	}

	err = Client.NewRequest(&req, &res)
	if err != nil {
		return 0, err
	}
	if resData.Ok != true {
		return 0, fmt.Errorf("not valida body req=%+v", resData)
	}

	return resData.Result.MessageID, err
}

func SendMessage(chatId int, text string, replyId int, mode string) error {

	m := MessageReq{
		ChatId: chatId,
		Text:   text,
		Mode:   mode,
	}

	if mode == "" {
		m.Mode = "Markdown"
	}

	if replyId != 0 {
		m.ReplyToMessageId = &replyId
	}

	req := requests.Request{
		Method:   http.MethodPost,
		Uri:      "sendMessage",
		JsonBody: &m,
		Flags: requests.RequestFlags{
			IsBodyString: true,
		},
	}
	res := requests.Response{}
	err := Client.NewRequest(&req, &res)
	return err
}
