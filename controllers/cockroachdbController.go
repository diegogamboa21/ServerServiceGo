package controllers

import (
	"database/sql"
	"log"

	"../models"

	_ "github.com/lib/pq"
)

var db *sql.DB

//InitDB create a connection with the database on the Amazon Web Services
func ConnectDB() {
	var err error
	connStr := "postgresql://root@Cockroach-ApiLoadB-7X3AGSGVMTI-1171392840.us-west-1.elb.amazonaws.com:26257?application_name=cockroach&sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

}

func InsertDomainOnDB(domain *models.Domain) {
	var IdDomain int64
	query, err := db.Prepare("INSERT INTO ServerService.Domain (ServersChanged, SSLGrade, PreviusSSLGrade, Logo, Title, IsDown) VALUES ($1,$2,$3,$4,$5,$6) RETURNING IdDomain;")
	if err != nil {
		log.Fatal("Insert Domain: ", err)
	}
	err = query.QueryRow(domain.ServersChanged, domain.SSLGrade, domain.PreviusSSLGrade, domain.Logo, domain.Title, domain.IsDown).Scan(&IdDomain)
	if err != nil {
		log.Fatal("Insert Domain QueryRow: ", err)
	}

	for _, s := range domain.Servers {
		InsertServerOnDB(IdDomain, s)
	}
	defer db.Close()
}

func InsertServerOnDB(IdDomain int64, server models.Server) {
	var IdServer int64
	//fmt.Println("InsertServerOnDB")
	query, err := db.Prepare("INSERT INTO ServerService.Server (Address, SSLGrade, Country, Owner, IdDomain) VALUES ($1,$2,$3,$4,$5);")
	if err != nil {
		log.Fatal("Insert Server: ", err)
	}

	query.QueryRow(server.Address, server.SSLGrade, server.Country, server.Owner, IdDomain).Scan(&IdServer)
}

func CalculateServersChanged(domain *models.Domain) {

}

func CalculatePreviusSSLGrade(domain *models.Domain) {

}
