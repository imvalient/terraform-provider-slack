package slack

import "github.com/nlopes/slack"

type client_api struct {
	token string
}

func (api *client_api) Client() (*slack.Client, error) {
	return slack.New(api.token), nil
}