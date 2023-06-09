import { createBrowserRouter, RouterProvider } from "react-router-dom";
import useLogin from "../hooks/useLogin";
import EditCategory from "../pages/EditCategory";
import AddPost from "../pages/AddPost";
import Home from "../pages/Home";
import Login from "../pages/Login";
import Post from "../pages/Post";
import EditPost from "../pages/EditPost";

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
        }, {
            path: "edit-category",
            element: <EditCategory />
        }, {
            path: 'edit-post/:postId',
            element: <EditPost />
        }]
    }
]


const MainRouter = () => {
    const isLog = useLogin();
    const router = isLog ? authRouterPath : routerPath;
    return <RouterProvider router={createBrowserRouter(router)} />
}

export default MainRouter;