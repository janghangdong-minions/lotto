package main

import (
	"encoding/json"
	"net/http"
)

func handleRandom(w http.ResponseWriter, r *http.Request) {
	nums := GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)
	data, err := json.Marshal(nums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
