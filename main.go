package main

import (
	"net/http"
	"os"
)

var data categorie
var Allitem AllItem
var api = "https://groupietrackers.herokuapp.com/api"

func Start(w http.ResponseWriter, r *http.Request) {
	var body, err = JsonOrder(api)
	JsonconvertApi(body, err)
	for _, category := range []string{data.Artists, data.Locations, data.Dates, data.Relation} {
		AssignData(body, err, category)
	}
}

func main() {
	if len(os.Args) == 1 {
		serv()
	}
}
