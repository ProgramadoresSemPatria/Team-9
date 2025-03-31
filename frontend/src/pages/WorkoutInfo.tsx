import { Fire, Timer } from 'phosphor-react';
import ExerciseInfo from '../components/ExerciseInfo';
import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';
import { Exercise, TrainingDay } from '../types';
import { useParams } from 'react-router';
import getExercises from '../services/exercises/getAll';
import getTrainingDayById from '../services/trainingDay/getById';

const WorkoutInfo = () => {
    const { id } = useParams();

    if (!id) return;

    const [trainingDay, setTrainingDay] = useState<TrainingDay>();
    const [exercises, setExercises] = useState<Exercise[]>([]);

    useEffect(() => {
        const token = Cookies.get('auth_token');

        if (!token) throw new Error('JWT token invalid');

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

    return (
        <div className="mx-auto flex min-h-screen max-w-4xl flex-col gap-5 p-8">
            <div className="flex flex-col gap-1.5">
                <h2 className="text-lg font-bold">{trainingDay?.title}</h2>
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
