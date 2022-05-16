package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var Token string

type Slack struct {
	Channel    string `yaml:"channel"`
	Color      string `yaml:"color"`      // Can either be one of good (green), warning (yellow), danger (red), or any hex color code (eg. #439FE0).
	Expression string `yaml:"expression"` // Expression to evaluate condition and send message.
	Message    string `yaml:"message"`
	Title      string `yaml:"title"`
}

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
func Send(channel, title, message, color string) error {
	if len(Token) == 0 {
		return errors.New("Plugin Slack: Please define token.")
	}

	msg := &Message{
		Text:    title,
		Channel: channel,
	}

	msg.addAttachment(&Attachment{
		Color: color,
		Text:  message,
	})

	err, _ := hook(msg)
	return err
}

func (m *Message) addAttachment(a *Attachment) {
	m.Attachments = append(m.Attachments, a)
}

func hook(msg *Message) (error, int) {
	jsonValues, _ := json.Marshal(msg)

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/"+Token,
		bytes.NewReader(jsonValues),
	)

	if err != nil {
		return err, 0
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err, 0
	}

	return nil, resp.StatusCode
}
