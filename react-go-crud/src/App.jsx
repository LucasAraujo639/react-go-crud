import React from 'react'

export const App = () => {
  return (
    <div>
      <h1> hello world</h1>
      <button onClick={async() => {
       const response = await fetch('http://localhost:3000/users')
       const data = await response.json()
       console.log(data)
       }}> Obtener datos </button> 
    </div>
  )
}
