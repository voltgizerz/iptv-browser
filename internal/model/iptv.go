package model

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Channel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	AltNames   []string `json:"alt_names"`
	Network    string   `json:"network"`
	Owners     []string `json:"owners"`
	Country    string   `json:"country"`
	Categories []string `json:"categories"`
	IsNSFW     bool     `json:"is_nsfw"`
	Launched   string   `json:"launched"`
	Closed     string   `json:"closed"`
	ReplacedBy string   `json:"replaced_by"`
	Website    string   `json:"website"`
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
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Country    string   `json:"country"`
	Categories []string `json:"categories"`
	Logo       string   `json:"logo"`
	IsNSFW     bool     `json:"is_nsfw"`
}
