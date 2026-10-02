import { useState } from "react"
import { Expense } from "./expense"
import { Input, type expense } from "./input"

const d: expense[] = [
  {amount: 10, currency: "INR", date: new Date().toISOString().split('T')[0] ?? "1970-01-01", uuid: "93c7850d-0c6a-49cd-8af8-b20c194e23f6"},
  {amount: 20, currency: "INR", date: new Date(Date.now() - 8.64e+7).toISOString().split('T')[0] ?? "1970-01-01", uuid: "eea73d01-47c6-4f47-9533-b3bc3fa15956"}
]

export function ExpensesTable() {

  const [data, setData] = useState<expense[]>(d)

  /*
   * function to get elements from golang server
   * the received elements should have a uuid, 
   * which is used in the Expense prop as id
  */
  
  const changeExpenseInArray = (exp: expense, id?: string) => {
    if(id === undefined) throw new Error(`id undefined when trying to change element ${exp}`);

    setData(prevData => prevData.map((e) => {
      return ((e.uuid === id) ? {...exp, uuid: id} : e)
    }))
  }

  return (
    <div>
      {data.map((exp) => {
        return ( 
          <div key={exp.uuid}>
          <Expense exp = {exp} />
          <Input id={exp.uuid} copyStateFn={changeExpenseInArray} />
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
