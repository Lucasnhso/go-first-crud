package main

import (
	"crud/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()

	fmt.Println("Running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}