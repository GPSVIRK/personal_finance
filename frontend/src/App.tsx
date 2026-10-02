import { updateSubmittedExpense } from "./api_go_server";
import { ExpensesTable } from "./expenses_table";
import { GetButton } from "./getButton";
import { Input } from "./input";

export function App() {
  return <div>
    <Input copyStateFn={updateSubmittedExpense} />
    <ExpensesTable />
    <GetButton />
  </div>;
}

