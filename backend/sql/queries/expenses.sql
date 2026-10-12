-- name: InsertExpense :one
INSERT INTO expenses(
  amount,
  currency,
  date_of_expense,
  updated_at,
  created_at
) VALUES(
  $1, $2, $3,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
)
RETURNING *;

-- name: GetExpenses :many
SELECT * FROM expenses;

-- name: UpdateExpense :one
UPDATE expenses SET amount=$1, currency=$2, date_of_expense=$3
WHERE id=$4
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id=$1;
