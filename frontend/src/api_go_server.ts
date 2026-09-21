import type { expense } from './input'

let submittedExpense: expense = { amount: 0, currency: "INR", date: new Date().toISOString().split('T')[0] ?? "1970-01-01" };

export let submittedExpenseJSON: string;

export const updateSubmittedExpense = (newExpense: expense) => {
  submittedExpense = newExpense;
  console.log("updated SubmittedExpense: ", submittedExpense);
};

/*
 * function to send request to go server and log messages
 * use JSON.stringify for converting the object to the correct thing
*/

/* 
 * this will handle connection with the go server established in backend
 * the way to do this is to connect to the port that it is established at 
 * then send the http requests, ig ill use handlers for that?
 * maybe in diff files but idk
 */
