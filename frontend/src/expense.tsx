import type { expense } from "./input";

export function Expense(props : { exp:expense}){
  const exp = props.exp

  return (
    <div>
      <p>spent {exp.currency} {exp.amount} at {exp.date} </p>
    </div>
  )
}
