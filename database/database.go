package database

import (
	"database/sql"
	"log"
	"time"
)

//DbConn database connection object
var DbConn *sql.DB

//SetupDatabase sets up database connection
func SetupDatabase() {
	var err error

	//needs mysql local instance user/password
	DbConn, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
