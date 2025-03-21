package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, mssg string){
	if code > 499 {
		log.Println("Responding with 5XX error: ",mssg)
	}

	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: mssg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to Marshal Json response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}