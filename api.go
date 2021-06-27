package main

import (
	"encoding/json"
	"net/http"
)

func handleRendom(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal("test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
