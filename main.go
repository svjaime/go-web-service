package main

import (
	"net/http"

	"github.com/svjaime/webapp/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)

	http.ListenAndServe(":8000", nil)
}
