import { axios } from '../utils';

const handleLogin = (name, password) => {
    return axios.post('/login', {
        name,
        password
    }, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}

const getAllCategory = () => {
    return axios.get('/all-category')
}

export {
    handleLogin,
    getAllCategory
}