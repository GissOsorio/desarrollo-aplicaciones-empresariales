import axios from 'axios';
import API_ROUTES from './apiRoutes';

const axiosInstance = axios.create({
    timeout: 10000,
});


const api = {
    get: (routeKey) => {
        console.log(`GET request to ${API_ROUTES[routeKey]}`);
        return axiosInstance.get(`http://localhost:8028/${routeKey}`);
    },
    post: (routeKey, data) => {
        console.log(`POST request to ${API_ROUTES[routeKey]}`);
        return axiosInstance.post(`http://localhost:8028/${routeKey}/`, data);
    },
    put: (routeKey, data) => {
        console.log(`PUT request to ${API_ROUTES[routeKey]}`);
        return axiosInstance.put(`http://localhost:8028/${routeKey}/`, data);
    },
    delete: (routeKey) => {
        console.log(`DELETE request to ${API_ROUTES[routeKey]}`);
        return axiosInstance.delete(`http://localhost:8028/${routeKey}/`);
    },
};

export default api;
