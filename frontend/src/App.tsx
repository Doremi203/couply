import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { AuthPage } from "./pages/AuthPage";
import { SplashPage } from "./pages/SplashPage";
import { HomePage } from "./pages/HomePage";
import { ProfilePage } from "./pages/ProfilePage";
import ChatPage from "./pages/ChatPage/components/ChatPage";
import { LikesPage } from "./pages/LikesPage";

const router = createBrowserRouter([
  {
    path: "auth",
    element: <AuthPage />,
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
    path: "chat",
    element: <ChatPage />,
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
