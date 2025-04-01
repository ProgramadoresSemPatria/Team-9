import { api } from '../../lib/axios';

const deleteExercise = async (exerciseId: string, token: string) => {
    try {
        const response = await api.delete(`/exercises/${exerciseId}`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default deleteExercise;
