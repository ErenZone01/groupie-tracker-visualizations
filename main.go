package main

import (
	"os"
)

var data categorie
var Allitem AllItem

func Start(){
	var api = "https://groupietrackers.herokuapp.com/api"
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
