import { api } from '../../lib/axios';

const getExercises = async (workoutDayId: string, token: string) => {
    try {
        const response = await api.get(`/workout-days/${workoutDayId}/exercises`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default getExercises;
