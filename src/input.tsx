import { useState } from 'react'

interface expense {
  amount: number,
  date: Date,
  currency: string
}

export function Input() {
  const [expense, setExpense] = useState<expense>({
    amount: 0,
    date: new Date(),
    currency: "INR"
  });

  const handleChangeAmt = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (typeof event.target.value === 'number' && !Number.isNaN(event.target.value)) {
      setExpense({
        ...expense,
        amount: event.target.value
      }); // this will leave the function immediately, cant run anything after it
    }
  }

  const handleChangeDate = (event: React.ChangeEvent<HTMLInputElement>) => {
    setExpense({
      ...expense,
      currency: event.target.value
    }); // this will leave the function immediately, cant run anything after it
  }

  const handleSelect = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setExpense({
      ...expense,
      currency: event.target.value
    })
  }

  return (<div><input
    type="number"
    onChange={handleChangeAmt}
    value={expense.amount} />
    <select value={expense.currency} onChange={handleSelect}>
      <option value="INR">INR</option>
    </select>
    <input
      type="date"
      value={expense.date.toDateString()}
      onChange={handleChangeDate} />
  </div>);
}

/* 
 * basically the handleChange detects any keystrokes and immediately puts in the text var
 * state is changed by setText, the state is of text 
 * putting in a div since it can only return one value
 *
 * the setter for the state thing will return immediately
 */
