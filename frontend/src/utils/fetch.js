import axios from "axios";
import env from "react-dotenv";

const tokenKey = env.AUTH_TOKEN;
const instance = axios.create({
    baseURL: env.REQUEST_URL
});

instance.defaults.headers[tokenKey] = sessionStorage.getItem(tokenKey);


export default instance;