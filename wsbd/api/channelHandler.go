package api

import (
	"net/http"
)

func GetChannelList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetChannel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
