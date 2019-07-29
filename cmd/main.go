package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/seryair/golang/internal/app/urlmanager"
)

func main() {

	http.HandleFunc("/", urlmanager.GetLocation)
	
	http.HandleFunc("/add", urlmanager.AddLocation)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
