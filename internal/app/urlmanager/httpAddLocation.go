package urlmanager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

func AddLocation(
	w http.ResponseWriter,
	r *http.Request) {

	if r.URL.Path != "/add" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		locationRequest := &LocationRequest{}
		locationRequest.UnmarshalJSON(bodyBytes)
		
		_, err1 := addLocation(*locationRequest)
		
		if err1 != nil {
			log.Fatal(err1)
		}

		fmt.Fprint(w, "Location Added successfully")
		return

	default:
		fmt.Fprintf(w, "Sorry, only POST method supported.")
	}
}
