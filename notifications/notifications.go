package notifications

import (
	"bytes"
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/lambda-engine/engine/env"
)

type Message struct {
	Username string `json:"username,omitempty"`
	Channel  string `json:"channel,omitempty"`
	Text     string `json:"text"`
	Pretext  string `json:"pretext,omitempty"`
}

func SimpleNotification(r *http.Request, text, pretext string) error {
	ctx := appengine.NewContext(r)
	return Notification(r, env.GetValue(ctx, "slack.webhook"), env.GetValue(ctx, "slack.channel"), env.GetValue(ctx, "slack.username"), text, pretext)
}

func Notification(r *http.Request, webhook, channel, username, text, pretext string) error {

	message := Message{
		username,
		channel,
		text,
		pretext,
	}
	m, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(m))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// post the request to Slack
	client := urlfetch.Client(appengine.NewContext(r))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
