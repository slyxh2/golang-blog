import { createBrowserRouter, RouterProvider } from "react-router-dom";
import useLogin from "../hooks/useLogin";
import Home from "../pages/Home/home";
import Login from "../pages/Login";

const routerPath = [
    {
        path: "/",
        element: <Login />
    }
]

const authRouterPath = [
    {
        path: "/",
        element: <Home />
    }
]


const MainRouter = () => {
    const isLog = useLogin();
    const router = isLog ? authRouterPath : routerPath;
    return <RouterProvider router={createBrowserRouter(router)} />
}

export default MainRouter;