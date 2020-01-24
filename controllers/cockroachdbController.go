package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"../models"

	//Connection with CockroachDB
	_ "github.com/lib/pq"
)

var db *sql.DB

//ConnectDB create a connection with the database CockroachDB on the server Amazon Web Services
func ConnectDB() {
	var err error
	connStr := "postgresql://root@Cockroach-ApiLoadB-7X3AGSGVMTI-1171392840.us-west-1.elb.amazonaws.com:26257?application_name=cockroach&sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

}

//DisconnectDB close the connection with the database
func DisconnectDB() {
	defer db.Close()
}

//InsertDomainOnDB use the struct domain and insert the info on the database
func InsertDomainOnDB(domain *models.Domain) {
	id, exists := FindDomainOnDB(domain.Title)
	if exists && id != 0 {
		fmt.Println("id encontrado: ", id)
		diference := CalculateTimeDiference(domain, id)
		fmt.Println("diference: ", diference)
		if diference {
			CalculateServersChanged(domain, id)
			CalculatePreviusSSLGrade(domain, id)
			InsertChanges(domain, id)
		}
	} else {
		var idDom int64
		query, err := db.Prepare("INSERT INTO ServerService.Domain (ServersChanged, SSLGrade, PreviusSSLGrade, Logo, Title, IsDown, Time) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING idDomain;")
		if err != nil {
			log.Fatal("Error insert Domain: ", err)
		}
		if err = query.QueryRow(domain.ServersChanged, domain.SSLGrade, domain.PreviusSSLGrade, domain.Logo, domain.Title, domain.IsDown, domain.LastQuery).Scan(&idDom); err != nil {
			log.Fatal("Error insert Domain QueryRow: ", err)
		}

		for _, s := range domain.Servers {
			InsertServerOnDB(idDom, s)
		}
		defer db.Close()
	}

}

//InsertServerOnDB insert on the database each server from the domain
func InsertServerOnDB(id int64, server models.Server) {
	var serverID int64
	//fmt.Println("InsertServerOnDB")
	query, err := db.Prepare("INSERT INTO ServerService.Server (Address, SSLGrade, Country, Owner, IdDomain) VALUES ($1,$2,$3,$4,$5);")
	if err != nil {
		log.Fatal("Insert Server: ", err)
	}

	query.QueryRow(server.Address, server.SSLGrade, server.Country, server.Owner, id).Scan(&serverID)
}

//CalculateServersChanged function
func CalculateServersChanged(domain *models.Domain, id int64) {

	query, err := db.Query("SELECT Address, SSLGrade, Country, Owner FROM ServerService.server WHERE IdDomain = $1 ;", id)
	if err != nil {
		log.Fatal("Error insert Domain 3: ", err)
	}
	query.Close()

	for query.Next() {
		var address, sslGrade, country, owner string
		if err := query.Scan(&address, &sslGrade, &country, &owner); err != nil {
			log.Fatal(err)
		}

		for _, s := range domain.Servers {
			if address == s.Address {
				if sslGrade == s.SSLGrade && country == s.Country && owner == s.Owner {
					domain.ServersChanged = true
					fmt.Println("ServersChanged: ", domain.ServersChanged)
				}
			}
		}
	}

}

//CalculatePreviusSSLGrade function
func CalculatePreviusSSLGrade(domain *models.Domain, id int64) {

	domain.PreviusSSLGrade = domain.SSLGrade
	CalculateServersGrade(domain)

}

//CalculateTimeDiference is a function that calculate if the diference to time is 1 hour or more
func CalculateTimeDiference(domain *models.Domain, id int64) bool {
	var previusGrade string
	var lastQuery time.Time

	query, err := db.Prepare("SELECT PreviusSSLGrade, time FROM ServerService.Domain WHERE IdDomain = $1")
	if err != nil {
		log.Fatal("Error insert Domain 2: ", err)
	}
	if err = query.QueryRow(id).Scan(&previusGrade, &lastQuery); err != nil {
		log.Fatal("Error insert Domain QueryRow 2: ", err)
	}

	timeNow := time.Now()
	timeDiference := timeNow.Sub(lastQuery)
	fmt.Println("timeDiference: ", timeDiference)
	if timeDiference.Hours() >= float64(1) {
		domain.LastQuery = timeNow
		fmt.Println("TimeNow: ", domain.LastQuery)
		return true
	}
	return false
}

//FindDomainOnDB find the domain and return id
func FindDomainOnDB(title string) (int64, bool) {
	fmt.Println("Title: ", title)
	row, err := db.Query("SELECT IdDomain FROM ServerService.Domain WHERE Title LIKE '" + title + "';")
	if err != nil {
		return -1, false
	}
	defer row.Close()
	var id int64
	row.Next()
	row.Scan(&id)
	return id, true
}

//InsertChanges update the last changes from domain
func InsertChanges(domain *models.Domain, id int64) {
	query, err := db.Prepare("UPDATE ServerService.Domain SET ServersChanged = $1, PreviusSSLGrade = $2, Time = $3 WHERE IdDomain = $4")
	if err != nil {
		log.Fatal("Error insert Domain 4: ", err)
	}
	query.QueryRow(domain.ServersChanged, domain.PreviusSSLGrade, domain.LastQuery, id)

}
