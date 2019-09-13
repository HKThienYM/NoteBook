package main

import (
	"log"
	"net/http"

	servicehandler "./service-handler"
)

//TODO: how to separate API logic, business logic and response format logic
func main() {
	http.HandleFunc("/postWithComments", servicehandler.Get(false))

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
