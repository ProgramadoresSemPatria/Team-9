import { api } from '../../lib/axios';
import { CreateFlow } from '../../types';

const createFlow = async (createFlowParams: CreateFlow, token: string) => {
    try {
        const response = await api.post(`/flows`, createFlowParams, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default createFlow;
