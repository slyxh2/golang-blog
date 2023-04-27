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

const createCategory = (name) => {
    return axios.post('/create-category', { name });
}

const getPosts = (page, size, category) => {
    return axios.get('/allpost', {
        params: {
            page,
            size,
            category
        }
    })
}

const deletePost = (id) => {
    return axios.delete('/delete-post', {
        params: {
            id
        }
    })
}

const getPost = (id) => {
    return axios.get('/getpost', {
        params: {
            id
        }
    })
}
const editPost = (obj) => { //file, header, category
    return axios.post('/edit-post', obj)
}
export {
    handleLogin,
    getAllCategory,
    uploadPost,
    deleteCategory,
    editCategoty,
    createCategory,
    getPosts,
    deletePost,
    getPost,
    editPost
}