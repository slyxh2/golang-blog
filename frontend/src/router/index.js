import { createBrowserRouter, RouterProvider } from "react-router-dom";
import useLogin from "../hooks/useLogin";
import AddPost from "../pages/AddPost";
import Home from "../pages/Home";
import Login from "../pages/Login";
import Post from "../pages/Post";

const routerPath = [
    {
        path: "/",
        element: <Login />
    }
]

const authRouterPath = [
    {
        path: "/",
        element: <Home />,
        children: [{
            path: "",
            element: <Post />
        }, {
            path: ":categoryId",
            element: <Post />
        }, {
            path: "add-post",
            element: <AddPost />
        }]
    }
]


const MainRouter = () => {
    const isLog = useLogin();
    const router = isLog ? authRouterPath : routerPath;
    return <RouterProvider router={createBrowserRouter(router)} />
}

export default MainRouter;