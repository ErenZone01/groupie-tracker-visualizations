package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func serv() {
	http.HandleFunc("/", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/icone/", http.StripPrefix("/icone/", http.FileServer(http.Dir("./icone/"))))
	fmt.Println("Serveur en cours d'exécution sur http://localhost:8080/")
	http.HandleFunc("/artist/", handlerPost)
	http.ListenAndServe(":8080", nil)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get(api)
	if err != nil {
		errorFile(w, r, "Error 500")
	} else {
		Allitem.Searchs = Search{}
		Start(w, r)
		url := strings.Split(r.URL.Path, "/")
		if len(url) == 3 {
			id := url[len(url)-1]
			Id, err := strconv.Atoi(id)
			if err != nil {
				var typeError = "Error 404"
				w.WriteHeader(http.StatusNotFound)
				errorFile(w, r, typeError)
			} else if Id <= 0 || Id > 52 {
				var typeError = "Error 404"
				w.WriteHeader(http.StatusNotFound)
				errorFile(w, r, typeError)
			} else {
				file, err := template.ParseFiles("template/details.html")
				if err != nil {
					var typeError = "Error 500"
					w.WriteHeader(http.StatusInternalServerError)
					errorFile(w, r, typeError)
				} else {
					Allitem.Id = Id - 1
					file.Execute(w, Allitem)
				}
			}

		} else {
			var typeError = "Error 404"
			w.WriteHeader(http.StatusNotFound)
			errorFile(w, r, typeError)
		}
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get(api)
	if err != nil {
		errorFile(w, r, "Error 500")
	} else {
		Allitem.Searchs = Search{}
		Start(w, r)
		if r.URL.Path != "/" {
			var typeError = "Error 404"
			w.WriteHeader(http.StatusNotFound)
			errorFile(w, r, typeError)
		} else {
			file, err := template.ParseFiles("template/index.html")
			if err != nil {
				var typeError = "Error 500"
				w.WriteHeader(http.StatusInternalServerError)
				errorFile(w, r, typeError)
			} else {
				file.Execute(w, Allitem)
			}
		}
	}

}
