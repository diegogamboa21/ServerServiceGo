package controllers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//InitDB create a connection with the database on the AWS
func InitDB() {

	connStr := "postgresql://root@Cockroach-ApiLoadB-7X3AGSGVMTI-1171392840.us-west-1.elb.amazonaws.com:26257?application_name=cockroach&sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		fmt.Println("No se pudo crear")
		log.Fatal(err)
	}
	/*
		if _, err := db.Exec(
			"INSERT INTO accounts (id, balance) VALUES (3, 1000), (4, 250)"); err != nil {
			log.Fatal(err)
		}
	*/
	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}

}
