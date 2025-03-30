import { useNavigate, useParams } from 'react-router';
import plusIcon from '../assets/plus.svg';
import TrainingDayContainer from '../components/TrainingDayContainer';
import { useEffect, useState } from 'react';
import getTrainingDayByFlowId from '../services/trainingDay/getByFlowId';
import Cookies from 'js-cookie';
import { Flow, TrainingDay } from '../types';
import getFlowById from '../services/flows/getById';

const FlowDetailsPage = () => {
    const [flow, setFlow] = useState<Flow>();
    const [daysOfTraining, setDaysOfTraining] = useState<TrainingDay[]>([]);

    const { id } = useParams();

    if (!id) return;

    const navigate = useNavigate();

    const handleClick = () => {
        navigate(`/add-new-day/${id}`);
    };

    useEffect(() => {
        const token = Cookies.get('auth_token');

        if (!token) throw new Error('JWT token invalid');

        const getFlow = async (flowId: string) => {
            const response = await getFlowById(flowId, token);

            if (response?.status === 200) {
                setFlow(response.data);
            }
        };

        const getDaysOfTraining = async (flowId: string) => {
            const response = await getTrainingDayByFlowId(flowId, token);

            if (response?.status === 200) {
                setDaysOfTraining(response.data);
            }
        };

        getDaysOfTraining(id);
        getFlow(id);
    }, []);

    return (
        <div className="flex w-full flex-col p-7 md:items-center">
            <div className="mb-4 flex w-full flex-col gap-2 md:items-center">
                <h1 className="text-2xl font-bold">{flow?.title}</h1>
                <h3 className="text-xl">{flow?.level}</h3>
            </div>
            <div className="flex flex-col items-center gap-4">
                {daysOfTraining.length > 0 ? (
                    daysOfTraining.map((trainingDay) => (
                        <TrainingDayContainer
                            key={trainingDay.id}
                            trainingDay={trainingDay}
                        />
                    ))
                ) : (
                    <span>No days of training added</span>
                )}
                <button
                    onClick={handleClick}
                    className="w-full cursor-pointer rounded-md border border-transparent bg-black px-2 py-2 text-xl text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:w-64"
                >
                    <div className="relative flex w-full items-center justify-center">
                        <img
                            src={plusIcon}
                            alt="plus"
                            className="absolute left-0 h-6 w-6"
                        />
                        <span className="text-lg">Add training day</span>
                    </div>
                </button>
            </div>
        </div>
    );
};

export default FlowDetailsPage;
