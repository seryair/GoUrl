package urlmanager

//easyjson:json
type LocationResponse struct {
	Location string `json:"location"`
}

//easyjson:json
type LocationRequest struct {
	Domain string `json:"domain"`
	Path   string `json:"path"`
}
