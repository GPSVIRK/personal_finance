-- name: InsertExpense :one
INSERT INTO expenses(
  amount,
  date_of_expense,
  updated_at,
  created_at
) VALUES(
  $1, $2,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
)
RETURNING *;

-- name: GetExpenses :many
SELECT * FROM expenses;
