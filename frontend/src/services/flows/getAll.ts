import { api } from '../../lib/axios';

const getFlowsByUser = async (token: string) => {
    try {
        const response = await api.get(`/flows`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default getFlowsByUser;
