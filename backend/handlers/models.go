package handlers

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/GPSVIRK/Personal_Finance_Manager/backend/internal/database"
)

type ApiConfig struct{
	DB *database.Queries
}

type Expense struct {
	ID            uuid.UUID `json:"id"`
	Amount        int		`json:"amount"`
	Currency      string	`json:"currency"`
	DateOfExpense time.Time `json:"date_of_expense"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

func databaseExpenseToExpense(dbExp database.Expense) (Expense, error){
	amt,err := strconv.Atoi(dbExp.Amount)
	if err!=nil{
		return Expense{}, err
	}

	return Expense{
		ID:            dbExp.ID,
		Amount:        amt,
		Currency:      dbExp.Currency,
		DateOfExpense: dbExp.DateOfExpense,
		UpdatedAt:     dbExp.UpdatedAt,
		CreatedAt:     dbExp.CreatedAt,
	}, nil
}

func databaseExpensesToExpenses(dbExps []database.Expense) ([]Expense, error){
	expenses := []Expense{}

	for _,dbExp := range dbExps{
		expense,err := databaseExpenseToExpense(dbExp)
		if err!=nil {
			return []Expense{}, err
		}
		expenses = append(expenses, expense)
	}

	return expenses,nil
}
