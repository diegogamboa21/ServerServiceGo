package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//Server only contain the instance http
type Server struct {
	server *http.Server
}

//NewServer create a connection http
func NewServer(mux *chi.Mux) *Server {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Server{s}
}

//Run permit to the server run online
func (s *Server) Run() {
	log.Fatal(s.server.ListenAndServe())
}

//Routes initiliazed the properties to server
func Routes() *chi.Mux {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/items", SearchListItems)
	mux.Get("/{domain}", SearchDomain)
	return mux
}

//SearchDomain accest to the database and insert the info about this domain
func SearchDomain(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "domain")

	domain := ReadDomainInfo(page)

	ConnectDB()
	InsertDomainOnDB(page, &domain)
	DisconnectDB()

	response, err := json.Marshal(domain)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

//SearchListItems find every element in the database
func SearchListItems(w http.ResponseWriter, r *http.Request) {

	ConnectDB()
	items := FindListItems()
	DisconnectDB()

	response, err := json.Marshal(items)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
