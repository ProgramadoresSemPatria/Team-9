import { api } from '../../lib/axios';
import { Login } from '../../types';

const login = async (loginParams: Login) => {
    try {
        const response = await api.post(`/login`, loginParams);

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default login;
