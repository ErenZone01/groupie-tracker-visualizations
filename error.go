package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Error struct {
	Texte string
}

func errorFile(w http.ResponseWriter, r *http.Request, Texte string) {
	file, err := template.ParseFiles("template/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print("Page inexistante")
		os.Exit(0)
	} else {
		var error Error
		error.Texte = Texte
		file.Execute(w, error)
	}

}
