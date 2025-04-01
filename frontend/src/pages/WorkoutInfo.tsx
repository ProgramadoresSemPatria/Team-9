import { Fire, Timer, Trash } from 'phosphor-react';
import ExerciseInfo from '../components/ExerciseInfo';
import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';
import { Exercise, TrainingDay } from '../types';
import { useNavigate, useParams } from 'react-router';
import getExercises from '../services/exercises/getAll';
import getTrainingDayById from '../services/trainingDay/getById';
import deleteTrainingDay from '../services/trainingDay/delete';

const WorkoutInfo = () => {
    const [trainingDay, setTrainingDay] = useState<TrainingDay>();
    const [exercises, setExercises] = useState<Exercise[]>([]);

    const navigate = useNavigate();

    const { id } = useParams();

    if (!id) return;

    const token = Cookies.get('auth_token');

    if (!token) throw new Error('JWT token invalid');

    useEffect(() => {
        const getTrainingDay = async () => {
            const response = await getTrainingDayById(id, token);

            if (response?.status !== 200) {
                throw new Error('Error to get flows');
            }

            setTrainingDay(response.data);
        };

        const getExercisesList = async () => {
            const response = await getExercises(id, token);

            if (response?.status !== 200) {
                throw new Error('Error to get flows');
            }

            setExercises(response.data);
        };
        getExercisesList();
        getTrainingDay();
    }, []);

    const handleDeleteWorkout = async () => {
        if (!trainingDay?.id) return;

        const response = await deleteTrainingDay(trainingDay.id, token);

        if (response?.status !== 200) {
            throw new Error('Error to delete training day');
        }

        navigate(-1);
    };

    return (
        <div className="mx-auto flex min-h-screen max-w-4xl flex-col gap-5 p-8">
            <div className="flex flex-col gap-1.5">
                <div className="flex items-center justify-between">
                    <h2 className="text-lg font-bold">{trainingDay?.title}</h2>
                    <button
                        onClick={handleDeleteWorkout}
                        className="cursor-pointer rounded-md bg-red-700 p-2 text-white transition duration-200 hover:bg-red-800"
                    >
                        <Trash size={20} />
                    </button>
                </div>
                <div className="flex gap-3">
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Fire weight="fill" className="text-gray-400" />
                            <p>
                                <span>{exercises.length}</span> exercises
                            </p>
                        </span>
                    </div>
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Timer weight="fill" className="text-gray-400" />
                            <p>
                                <span>{trainingDay?.duration}</span> min
                            </p>
                        </span>
                    </div>
                </div>
            </div>
            <div className="flex flex-col gap-5">
                {exercises.map((exercise) => (
                    <ExerciseInfo key={exercise.id} exerciseInfo={exercise} />
                ))}
            </div>
        </div>
    );
};

export default WorkoutInfo;
