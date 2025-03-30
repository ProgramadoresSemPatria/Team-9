import { api } from '../../lib/axios';

const getTrainingDayByFlowId = async (flowId: string, token: string) => {
    try {
        const response = await api.get(`/flows/${flowId}/workout-days`, {
            withCredentials: true,
            headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default getTrainingDayByFlowId;
