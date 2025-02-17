package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestData struct {
	Task string `json:"task"`
}

type ResponseData struct {
	Task string `json:"task"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	response := ResponseData{
		Task: data.Task + "LAME SHIT",
	}

	err_ := json.NewEncoder(w).Encode(response)

	if err_ != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", helloHandler).Methods("GET")
	router.HandleFunc("/api/post", postHandler).Methods("POST")

	http.ListenAndServe("localhost:8080", router)
}
