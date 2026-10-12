package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/GPSVIRK/Personal_Finance_Manager/backend/internal/database"
)

func (apicfg *ApiConfig) GetExpenses(w http.ResponseWriter, r *http.Request){
	dbExpns, err := apicfg.DB.GetExpenses(r.Context())
	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Err getting expenses: %v",err))
		return
	}

	expns, err := databaseExpensesToExpenses(dbExpns)
	if err!=nil {
		respondWithError(w, 500, fmt.Sprintf("Err bad data: %v", err))
	}

	respondWithJSON(w, 200, expns)
}

func (apicfg *ApiConfig)PostExpenses(w http.ResponseWriter, r *http.Request){
	type params struct {
		Amount 			string 		`json:"amount"`
		Currency		string		`json:"currency"`
		DateOfExpense 	time.Time 	`json:"date_of_expense"`
	}

	decoder := json.NewDecoder(r.Body)

	p := params{}
	err := decoder.Decode(p)
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("Err decoding input: %v",err))
		return
	}

	dbExp, err := apicfg.DB.InsertExpense(r.Context(), database.InsertExpenseParams{
		Amount: p.Amount,
		Currency: p.Currency,
		DateOfExpense: p.DateOfExpense,
	})
	if err!=nil {
		respondWithError(w, 500, fmt.Sprintf("Err inserting expense: %v",err))
		return
	}

	exp, err := databaseExpenseToExpense(dbExp)
	if err!=nil{
		respondWithError(w,500,fmt.Sprintf("Err converting expense types: %v",err))
		return
	}

	respondWithJSON(w, 201, exp)
}

func (apicfg *ApiConfig)PatchExpenses(w http.ResponseWriter, r *http.Request){
	type params struct {
		Amount        string	`json:"amount"`
		Currency      string	`json:"currency"`
		DateOfExpense time.Time	`json:"date_of_expense"`
		ID            uuid.UUID	`json:"id"`
	}

	decoder := json.NewDecoder(r.Body)

	p:=params{}

	err := decoder.Decode(p)
	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("Err failed to decode for update: %v",err))
		return
	}

	dbExp, err := apicfg.DB.UpdateExpense(r.Context(), database.UpdateExpenseParams{
		Amount: p.Amount,
		Currency: p.Currency,
		DateOfExpense: p.DateOfExpense,
		ID: p.ID,
	})
	if err!=nil {
		respondWithError(w, 500, fmt.Sprintf("Err failed to update expense: %v", err))
		return
	}

	exp, err := databaseExpenseToExpense(dbExp)
	if err!=nil{
		respondWithError(w,500,fmt.Sprintf("Err failed to convert expense types: %v",err))
		return
	}

	respondWithJSON(w, 200, exp)
}

func (apiCfg *ApiConfig) DeleteExpense(w http.ResponseWriter, r *http.Request){
	type params struct {
		ID uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	p:=params{}

	err := decoder.Decode(p)
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("Err failed to decode: %v", err))
		return
	}

	err = apiCfg.DB.DeleteExpense(r.Context(), p.ID)
	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Err failed to delete expense: %v", err))
		return
	}

	var empty interface{}
	respondWithJSON(w, 200, empty)
}
