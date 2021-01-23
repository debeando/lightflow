package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var Token string

type Message struct {
	Text        string        `json:"text"`
	Channel     string        `json:"channel,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Color string `json:"color,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

// Retry is a method to call many times until satify return value.
func Send(channel, title, message, color string) {
	msg := &Message{
		Text:    title,
		Channel: channel,
	}

	msg.addAttachment(&Attachment{
		Color: color,
		Text:  message,
	})

	hook(msg)
}

func (m *Message) addAttachment(a *Attachment) {
	m.Attachments = append(m.Attachments, a)
}

func hook(msg *Message) int {
	jsonValues, _ := json.Marshal(msg)

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/"+Token,
		bytes.NewReader(jsonValues),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err)
	}

	return resp.StatusCode
}
