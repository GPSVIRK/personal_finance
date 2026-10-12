package handlers

import "net/http"

func Test(w http.ResponseWriter, r *http.Request){
	type testResponse struct{
		Message string `json:"message"`
	}
	respondWithJSON(w, 200, testResponse{
		Message: "TEST MESSAGE WOW",
	})
}
