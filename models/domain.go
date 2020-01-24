package models

import "time"

//Domain is the struct tha contain the webpage data
type Domain struct {
	Servers         []Server `json:"endpoints"`
	ServersChanged  bool
	SSLGrade        string
	PreviusSSLGrade string
	Logo            string
	Title           string
	IsDown          bool
	LastQuery       time.Time
}
