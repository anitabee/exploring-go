package server

import (
	"log"
	"net/http"
)

func WebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello server!</h1>"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
