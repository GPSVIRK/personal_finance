package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, errMsg string){
	if code>499{
		log.Println("Responding with 5XX level error: ", errMsg)
	}
	type errorResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: errMsg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)

	if err!=nil {
		log.Printf("Failed to marshal payload: %v",payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
