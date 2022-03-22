package main

import (
	"awesomeProject/racehub/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		http.Error(w, "cannot get threads", http.StatusInternalServerError)
	} else {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	}
}
