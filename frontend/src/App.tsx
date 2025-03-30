import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { AuthPage } from "./pages/AuthPage";
import { SplashPage } from "./pages/SplashPage";
import { HomePage } from "./pages/HomePage";
import { ProfilePage } from "./pages/ProfilePage";
import { LikesPage } from "./pages/LikesPage";
import { EnterInfoPage } from "./pages/EnterInfoPage";

const router = createBrowserRouter([
  {
    path: "auth",
    element: <AuthPage />,
  },
  {
    path: "enterInfo",
    element: <EnterInfoPage />,
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
  {
    path: "likes",
    element: <LikesPage />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
