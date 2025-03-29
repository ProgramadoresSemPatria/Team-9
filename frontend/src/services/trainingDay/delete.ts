import { api } from '../../lib/axios';

const deleteTrainingDay = async (trainingDayId: string, token: string) => {
    try {
        const response = await api.get(`/workout-days/${trainingDayId}`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default deleteTrainingDay;
