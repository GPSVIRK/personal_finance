import { useState } from 'react'

export function Input() {
  const [text, setText] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value); // this will leave the function immediately, cant run anything after it
  }

  return (<div><input
    type="text"
    onChange={handleChange}
    value={text} /></div>);
}

/* 
 * basically the handleChange detects any keystrokes and immediately puts in the text var
 * state is changed by setText, the state is of text 
 * i put it in a div for no reason
 */
