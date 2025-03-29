import { api } from '../../lib/axios';

const getTrainingDayByFlowId = async (flowId: string, token: string) => {
    try {
        const response = await api.get(`/workout-days/${flowId}`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default getTrainingDayByFlowId;
