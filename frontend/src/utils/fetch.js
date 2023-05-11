import axios from "axios";

const tokenKey = 'Authorization';
const instance = axios.create({
    baseURL: 'https://golang-blog-production.up.railway.app'
});

instance.defaults.headers[tokenKey] = sessionStorage.getItem(tokenKey);
instance.defaults.headers['Content-Type'] = 'multipart/form-data';

export default instance;