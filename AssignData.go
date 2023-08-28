package main

import (
	"fmt"
	"io"
	"net/http"
)

func JsonOrder(api string) ([]byte, error) {
	response, err := http.Get(api)
	if err != nil {
		serv()
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse de l'API :", err)

	}
	return body, err
}

func AssignData(body []byte, err error, datas string) {
	if data.Artists == datas {
		body, err = JsonOrder(datas)
		JsonconvertArtists(body, err)
	} else if data.Locations == datas {
		body, err = JsonOrder(datas)
		JsonconvertLocations(body, err)
	} else if data.Dates == datas {
		body, err = JsonOrder(datas)
		JsonconvertDates(body, err)
	} else if data.Relation == datas {
		body, err = JsonOrder(datas)
		JsonConvertRelation(body, err)
	}
}
