import { useState } from "react";
import { testConnection } from "./api_go_server";

export function DisplayThing(props:{text: string}){
  return <p> {props.text} </p>
}

export function GetButton(){
  const [display, setDisplay] = useState(false)
  const [thing, setThing] = useState("")
  const handleClick = async () =>{
    try{
      const res=await testConnection()
      setThing(res.message);
      setDisplay(true)

      setTimeout(() => {
        setDisplay(false)
      }, 899)
    }
    catch(error){
      console.log("error",error)
    }
  }
  return <div>
    <button onClick={handleClick}>TEST</button>
    {display && <DisplayThing text={thing} />}
  </div>
}
