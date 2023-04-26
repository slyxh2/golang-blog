import { axios } from '../utils';

const handleLogin = (name, password) => {
    return axios.post('/login', {
        name,
        password
    })
}

const getAllCategory = () => {
    return axios.get('/all-category')
}

const uploadPost = (file, header, categoryId) => {
    return axios.post('/upload', {
        file,
        header,
        categoryId
    })
}

export {
    handleLogin,
    getAllCategory,
    uploadPost
}