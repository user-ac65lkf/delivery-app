package pushV1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bakhtiyor-y/pushservice/internal/models"
	"io"
	"log"
	"net/http"
)

func SendTelegramMessage(message models.MessageDetails) {

	msg := &models.Message{
		ChatID: 634926850,
		Text:   message,
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println("telegram message err", err)
	}
	response, err := http.Post(
		"https://api.telegram.org/bot6547743868:AAHN062AH3te6ZMhIIS0xBiw-sIajdZ3g5E/sendMessage",
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Println("telegram message err", err)
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close telegram response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("failed to send successful request. Status was %q", response.Status))
	}
}
