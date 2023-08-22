import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {useAuth0} from "@auth0/auth0-react";
import LoginButton from "./components/login-button.tsx";
import Profile from "./components/profile.tsx";
import LogoutButton from "./components/logout-button.tsx";

function App() {
    const {isAuthenticated} =useAuth0()
  return (
    <>
<h1>Aplicaciones empresariales</h1>
        {!isAuthenticated ?  <LoginButton/> : null}
        {isAuthenticated ?   <LogoutButton/> : null}
        {isAuthenticated ?    <Profile/>: null}




    </>
  )
}

export default App
