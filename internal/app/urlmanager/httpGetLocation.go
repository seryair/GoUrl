package urlmanager

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/sync/singleflight"
)

var requestGroup singleflight.Group

func GetLocation(
	w http.ResponseWriter,
	r *http.Request) {

	if r.URL.Path != "/" {
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

		v, err, shared := requestGroup.Do(string(bodyBytes), func() (interface{}, error) {
			return getLocation(*locationRequest)
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		locationResponse := &LocationResponse{}

		if str, ok := v.(string); ok {
			locationResponse.Location = str
			if shared {
				log.Println("shared result")
			} else {
				log.Println("not a shared result")
			}
		} else {
			locationResponse.Location = ""
		}

		var mrs []byte
		mrs, errMrs := locationResponse.MarshalJSON()
		if errMrs != nil {
			log.Println(errMrs)
		}

		fmt.Fprint(w, string(mrs))
		return

	default:
		fmt.Fprintf(w, "Sorry, only POST method supported.")
	}
}
