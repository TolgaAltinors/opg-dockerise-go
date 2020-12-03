package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func home(w http.ResponseWriter, req *http.Request) {

	status := Status{Code: 200, Message: "Home"}

	response, err := json.Marshal(status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func status(w http.ResponseWriter, req *http.Request) {

	statusFromEnv := os.Getenv("APP_STATUS")

	status := Status{Code: 500, Message: statusFromEnv}

	if statusFromEnv == "OK" {
		status.Code = 200
	}

	response, err := json.Marshal(status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

// Additional endpoint to dynamically alter environment variable
func setStatus(w http.ResponseWriter, req *http.Request) {

	key, _ := req.URL.Query()["key"]

	os.Setenv("APP_STATUS", key[0])

	status := Status{Code: 200, Message: "APP_STATUS SET"}

	response, err := json.Marshal(status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/status/", status)
	http.HandleFunc("/setStatus/", setStatus)
	http.ListenAndServe(":8080", nil)
}
