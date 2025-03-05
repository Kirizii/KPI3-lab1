package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// При зверенні на порт 8795 з endpoint /time, повертаємо нинішній час у вигляді YYYY-MM-DDTHH:MM:SSZ
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}

//CBA
