import axios from 'axios'

const api = axios.create({
    baseURL: 'https://deezer-rank.herokuapp.com/'
});

export default api;

