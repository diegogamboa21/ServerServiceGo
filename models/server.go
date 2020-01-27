package models

//Server is the struct to define a server from Domain
type Server struct {
	Address  string `json:"address"`
	SSLGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}
