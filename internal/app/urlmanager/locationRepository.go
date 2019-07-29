package urlmanager

import (
	"fmt"
)

var LocationMapper = map[string][]string{
	"ynet.co.il": {"/home/0,7340,L-8,00.html", "/home/abc.html"},
	"sport5":     {"/home/123/ac.html", "/fds.html"},
	"walla":      {"/home/12.html"},
}

func getLocation(lR LocationRequest) (string, error) {

	v, ok := LocationMapper[lR.Domain]
	fmt.Println("The value:", v, "Present?", ok)

	if ok {
		if stringInSlice(lR.Path, v) {
			return "https://" + lR.Domain + lR.Path, nil
		}
	}
	return "", nil
}

func addLocation(lR LocationRequest) (string, error) {

	v, ok := LocationMapper[lR.Domain]
	fmt.Println("The value:", v, "Present?", ok)

	if ok {
	    v = append(v, lR.Path)
		LocationMapper[lR.Domain] = v
	} else {
		var paths []string
		paths = append(paths, lR.Path)
		LocationMapper[lR.Domain] = paths
	}
	return "", nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
