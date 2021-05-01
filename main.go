package main

import (
	"fmt"
	"go-webservice/product"
	"net/http"
)

const apiBasePath = "/api"

func main() {
	fmt.Println("api started")

	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
