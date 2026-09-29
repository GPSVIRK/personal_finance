import { Expense } from "./expense"
import type { expense } from "./input"

export function ExpensesTable() {
  const data: expense[]  = [
    {amount: 10, currency: "INR", date:new Date().toISOString().split('T')[0] ?? "1970-01-01"}
  ]

  /*
   * function to get elements from golang server
   * the received elements should have a uuid, 
   * which is used in the Expense prop as id
  */

  return (
    <div>
      {data.map((exp) => {
        return ( 
          <div>
          <Expense exp = {exp} />
          </div>
        )
      })}
    </div>
  )
}

/*
 * this will contain the values received from go,
 * there will be the ability to change the things
 * i think i should dynamically create components
 * for each expense, and the input is the expenses interface
*/
