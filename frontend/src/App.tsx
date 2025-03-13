import './App.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import AuthPage from './pages/AuthPage/AuthPage'
import SplashPage from './pages/SplashPage/SplashPage'
import HomePage from './pages/HomePage/HomePage'
import ProfilePage from './pages/ProfilePage/ProfilePage'


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
  {
    path: "profile",
    element: <ProfilePage />,
  },
])


function App() {
 

  return (
    <RouterProvider router={router}/>
);
}

export default App
