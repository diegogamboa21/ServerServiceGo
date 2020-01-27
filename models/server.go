package models

//Server is the struct to define a server from Domain
type Server struct {
	Address  string `json:"ipAddress"`
	SSLGrade string `json:"grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}
