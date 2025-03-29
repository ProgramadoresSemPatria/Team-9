import { api } from '../../lib/axios';
import { RegisterUser } from '../../types';

const registerUser = async (registerUserParams: RegisterUser) => {
    try {
        const response = await api.post(`/register`, registerUserParams);

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default registerUser;
