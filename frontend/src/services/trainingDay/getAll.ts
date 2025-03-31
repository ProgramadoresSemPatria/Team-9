import { api } from '../../lib/axios';

const getTrainingDays = async (token: string) => {
    try {
        const response = await api.get(`/workout-days`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default getTrainingDays;
