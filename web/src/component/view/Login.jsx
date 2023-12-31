import React from 'react'
import {Link} from "react-router-dom";
import { useNavigate } from 'react-router-dom';
import {UserContext} from '../../App'
// dispatch({type:"USER",payload:data.user})
//     const {state,dispatch} = useContext(UserContext)
function Login(){
  const {state,dispatch} = React.useContext(UserContext)
  const navigate = useNavigate()
  const [email,setEmail] = React.useState("")
  const [password,setPassword] = React.useState("")
  const url = "http://localhost:8887/auth/login"
  const data = {"email":"yale918@gmail.com","password":"12345"}
  const postData = (url, data) =>{
    fetch(url,{
      body: JSON.stringify(data),
      headers: {
        'content-type': 'application/json'
      },
      method: 'POST'
    })
    .then(res=>res.json())
    .then(data=>{
      console.log("data:")
      console.log(data)
      localStorage.setItem("user", data.name)
      localStorage.setItem("token", data.token)
      dispatch({type:"USER", payload:data.name})
      navigate("/chat")
    })
  }
  
  return (

    <div className="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-lg text-center">
        <h1 className="text-2xl font-bold sm:text-3xl">Login</h1>
        <br/>
        <h2 
          className="underline cursor-pointer"
          onClick={()=>{
            postData(url,data)
          }}  
        >Bypass Login</h2>
      </div>

      <form action="" className="mx-auto mt-8 mb-0 max-w-md space-y-4">
        <div>
          <label htmlFor="email" className="sr-only">Email</label>

          <div className="relative">
            <input
              type="email"
              className="w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm"
              placeholder="Enter email"
              onChange={(e)=>{
                setEmail(e.target.value)
              }}
            />

            <span className="absolute inset-y-0 right-4 inline-flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5 text-gray-400"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
                />
              </svg>
            </span>
          </div>
        </div>

        <div>
          <label htmlFor="password" className="sr-only">Password</label>
          <div className="relative">
            <input
              type="password"
              className="w-full rounded-lg border-gray-200 p-4 pr-12 text-sm shadow-sm"
              placeholder="Enter password"
              onChange={(e)=>{
                setPassword(e.target.value)
              }}
            />

            <span className="absolute inset-y-0 right-4 inline-flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5 text-gray-400"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                />
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                />
              </svg>
            </span>
          </div>
        </div>

        <div className="flex items-center justify-between">
          <p className="text-sm text-gray-500">
            No account?
            <Link className="underline" to="/regist">Sign up</Link>
            {/* <a className="underline" href="">Sign up</a> */}
          </p>

          <button
            type="submit"
            className="ml-3 inline-block rounded-lg bg-blue-500 px-5 py-3 text-sm font-medium text-white"

            onClick={(e)=>{
              e.preventDefault()
              postData(url,{email,password})
            }}
          >
            Sign in
          </button>
        </div>
      </form>
    </div>

  )
}

export default Login