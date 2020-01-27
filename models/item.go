package models

//Item struct
type Item struct {
	WebURL string `json:"WebURL"`
	Site   Domain `json:"info"`
}
