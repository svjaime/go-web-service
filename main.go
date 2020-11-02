package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/svjaime/webapp/database"
	"github.com/svjaime/webapp/product"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
