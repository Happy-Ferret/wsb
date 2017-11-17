package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListenAndServe(port string) {
	r := mux.NewRouter()

	s := r.PathPrefix("/channels").Subrouter()
	s.HandleFunc("", GetChannelList).Methods("GET")
	s.HandleFunc("", CreateChannel).Methods("POST")
	s.HandleFunc("/{key}", GetChannel)

	log.Fatal(http.ListenAndServe(port, r))
}
