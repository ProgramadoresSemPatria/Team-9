import { api } from '../../lib/axios';
import { CreateTrainingDay } from '../../types';

const createTrainingDay = async (
    createTrainingDayParams: CreateTrainingDay,
    token: string
) => {
    try {
        const response = await api.post(`/workout-days`, createTrainingDayParams, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default createTrainingDay;
