package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JsonconvertApi(body []byte, err error) {
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation de la réponse JSON :", err)
		return
	}
}

func JsonconvertArtists(body []byte, err error) {
	var a []artists // Créer une nouvelle instance de la structure artists
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation de la réponse JSON :", err)
		return
	}

	Allitem.Artists = a
}

func JsonconvertLocations(body []byte, err error) {
	var l = IndexLocation{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation de la réponse JSON :", err)
		return
	}
	for _, data := range l.Index {
		Allitem.Location = append(Allitem.Location, data)
	}
}

func JsonconvertDates(body []byte, err error) {
	var d = IndexDates{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation de la réponse JSON :", err)
		return
	}

	for _, data := range d.Index {
		for i, v := range data.Dates {
			data.Dates[i] = strings.TrimPrefix(v, "*")
		}
		Allitem.Dates = append(Allitem.Dates, data)
	}

}

func JsonConvertRelation(body []byte, err error) {
	var r = IndexRel{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation de la réponse JSON :", err)
		return
	}

	for _, valeurs := range r.Index {
		var data = RelaTab{}
		for cle, valeur := range valeurs.DatesLocations {
			var texte string = ""
			for i, v := range valeur {
				texte += v
				if i != len(valeur)-1 {
					texte += ", "
				}
			}
			data.DatesLocations = append(data.DatesLocations, cle+" : "+texte)
		}
		Allitem.Relations = append(Allitem.Relations, data)
	}

}
