package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GPSVIRK/Personal_Finance_Manager/backend/internal/database"
)

func GetExpenses(w http.ResponseWriter, r *http.Request){
	;
}

func (apicfg *ApiConfig)PostExpenses(w http.ResponseWriter, r *http.Request){
	type params struct {
		Amount 			string 		`json:"amount"`
		DateOfExpense 	time.Time 	`json:"date_of_expense"`
	}

	decoder := json.NewDecoder(r.Body)

	p := params{}
	err := decoder.Decode(p)
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("Err decoding input: %v",err))
		return
	}

	dbExpns, err := apicfg.DB.InsertExpense(r.Context(), database.InsertExpenseParams{
		Amount: p.Amount,
		DateOfExpense: p.DateOfExpense,
	})
	if err!=nil {
		respondWithError(w, 500, fmt.Sprintf("Err inserting expense: %v",err))
		return
	}

	exp, err := databaseExpenseToExpense(dbExpns)

	respondWithJSON(w, 201, exp)
}

func PatchExpenses(w http.ResponseWriter, r *http.Request){
	;
}
