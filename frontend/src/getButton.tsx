import { testConnection } from "./api_go_server";


export function GetButton(){
  const handleClick = async () =>{
    try{
      const res=await testConnection()
      console.log(res.message)
    }
    catch(error){
      console.log("error",error)
    }
  }
  return <button onClick={handleClick}>TEST</button>
}
