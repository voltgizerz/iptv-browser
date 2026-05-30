package model

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Channel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Stream struct {
	Channel string `json:"channel"`
	URL     string `json:"url"`
}

type Logo struct {
	Channel string `json:"channel"`
	URL     string `json:"url"`
}

type ChannelResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
}