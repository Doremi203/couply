import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import AuthPage from './pages/AuthPage/AuthPage'
import SplashPage from './pages/SplashPage/SplashPage'
import HomePage from './pages/HomePage/HomePage'


const router = createBrowserRouter([
  {
      path: "auth",
      element: <AuthPage/>,
  },
  {
      path: "/",
      element: <SplashPage />,
  },
  {
    path: "home",
    element: <HomePage />,
},
])


function App() {
 

  return (
    <RouterProvider router={router}/>
);
}

export default App
