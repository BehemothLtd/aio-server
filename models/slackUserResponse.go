package models

type SlackUserReponse struct {
	Ok   bool      `json:"ok"`
	User SlackUser `json:"user"`
}

type SlackUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
