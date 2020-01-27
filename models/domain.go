package models

import "time"

//Domain is the struct tha contain the webpage data
type Domain struct {
	Servers         []Server  `json:"endpoints"`
	ServersChanged  bool      `json:"servers_changed"`
	SSLGrade        string    `json:"grade"`
	PreviusSSLGrade string    `json:"previous_ssl_grade"`
	Logo            string    `json:"logo"`
	Title           string    `json:"title"`
	IsDown          bool      `json:"is_down"`
	LastQuery       time.Time `json:"time"`
}
