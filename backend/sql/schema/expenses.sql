CREATE TABLE expenses(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  amount DECIMAL(10,2) NOT NULL,
  currency CHAR(3) NOT NULL,
  date_of_expense TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL
);
