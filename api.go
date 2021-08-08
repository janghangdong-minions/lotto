package main

import (
	"encoding/json"
	"net/http"
)

func GetRandomnum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "이함수는 GET메소드만 지원합니다.", http.StatusBadRequest)
		return
	}
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
