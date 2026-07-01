package handlers

import (
	"html/template"
	"net/http"
)

func Music(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/music"{
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet{
		w.Header().Set("Allow", http.MethodGet)

		http.Error(w, "Method not allowed", 405)
		return
	}

	temp, err := template.ParseFiles("templates/music.html")
	if err != nil{
		http.Error(w, "Internal server error", 405)
		return
	}

	temp.Execute(w, nil)
}
