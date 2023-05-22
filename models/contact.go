package models

type Contact struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	Organization string `json:"organization"`
	Dolzhnost    string `json:"dolzhnost"`
	Mobile       string `json:"mobile"`
	UserId       int    `json:"user_id"`
}
