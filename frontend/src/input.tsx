import { useState } from 'react'
import { updateSubmittedExpense } from './api_go_server';

export interface expense {
  amount: number,
  date: string,
  currency: string
}

export function Input() {
  const [expense, setExpense] = useState<expense>({
    amount: 0,
    date: new Date().toISOString().split('T')[0] ?? "1970-01-01",
    currency: "INR"
  });

  const handleChangeAmt = (event: React.ChangeEvent<HTMLInputElement>) => {
    const inpAmnt = parseFloat(event.target.value);
    if (typeof inpAmnt === 'number') {
      setExpense({
        ...expense,
        amount: inpAmnt
      }); 
    }
  }

  const handleChangeDate = (event: React.ChangeEvent<HTMLInputElement>) => {
    setExpense({
      ...expense,
      date: event.target.value
    }); 
  }

  const handleSelect = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setExpense({
      ...expense,
      currency: event.target.value
    })
  }

  const handleSubmit = () => {
    updateSubmittedExpense(expense);
  }

  return (<div>
    <input
      type="number"
      onChange={handleChangeAmt}
      value={expense.amount} />
    <select value={expense.currency} onChange={handleSelect}>
      <option value="INR">INR</option>
    </select>
    <input
      type="date"
      value={expense.date}
      onChange={handleChangeDate} />
    <button onClick={handleSubmit}>Submit</button>
  </div>);
}

/* 
 * basically the handleChange detects any keystrokes and immediately puts in the text var
 * state is changed by setText, the state is of text 
 * putting in a div since it can only return one value
 *
 * the setter for the state thing will return immediately
 */
