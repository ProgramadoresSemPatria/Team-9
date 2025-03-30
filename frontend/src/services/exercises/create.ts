import { api } from '../../lib/axios';
import { CreateExercise } from '../../types';

const createExercise = async (
    createExerciseParams: CreateExercise,
    trainingDayId: string,
    token: string
) => {
    try {
        const response = await api.post(
            `/workout-days/${trainingDayId}/exercises`,
            createExerciseParams,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            }
        );

        return response;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};

export default createExercise;
