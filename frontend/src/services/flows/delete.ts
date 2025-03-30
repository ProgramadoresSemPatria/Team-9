import { api } from '../../lib/axios';

const deleteFlow = async (flowId: string, token: string) => {
    try {
        const response = await api.get(`/flows/${flowId}`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default deleteFlow;
