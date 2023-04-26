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

const deleteCategory = (id) => {
    return axios.delete('/delete-category', {
        params: { id }
    })
}

const editCategoty = (id, name) => {
    return axios.post('/edit-category', {
        id,
        name
    })
}

export {
    handleLogin,
    getAllCategory,
    uploadPost,
    deleteCategory,
    editCategoty
}