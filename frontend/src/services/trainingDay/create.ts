import { api } from '../../lib/axios';
import { CreateTrainingDay } from '../../types';

const createTrainingDay = async (
    createTrainingDayParams: CreateTrainingDay,
    flowId: string,
    token: string
) => {
    try {
        console.log(typeof createTrainingDayParams.duration);
        const response = await api.post(
            `/flows/${flowId}/workout-days`,
            createTrainingDayParams,
            {
                withCredentials: true,
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

export default createTrainingDay;
