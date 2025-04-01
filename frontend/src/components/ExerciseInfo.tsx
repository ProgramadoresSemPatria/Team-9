import { Barbell, Trash } from 'phosphor-react';
import deleteExercise from '../services/exercises/delete';
import Cookies from 'js-cookie';

interface ExerciseInfoProps {
    exerciseInfo: {
        id: string;
        title: string;
        sets: number;
        repetitions: number;
    };
}

const ExerciseInfo = ({ exerciseInfo }: ExerciseInfoProps) => {
    const handleDeleteExercise = async () => {
        if (!exerciseInfo.id) return;

        const token = Cookies.get('auth_token');

        if (!token) throw new Error('JWT token invalid');

        const response = await deleteExercise(exerciseInfo.id, token);

        if (response?.status !== 200) {
            throw new Error('Error to delete training day');
        }

        window.location.reload();
    };

    return (
        <div className="rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5">
            <div className="flex items-center justify-between rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-4">
                <div className="flex items-center gap-3 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-4">
                    <div className="rounded-md bg-zinc-200 p-3">
                        <Barbell
                            size={20}
                            weight="fill"
                            className="text-purple-500"
                        />
                    </div>
                    <div>
                        <h3 className="font-medium text-black">
                            {exerciseInfo.title}
                        </h3>
                        <div className="flex gap-2">
                            <p className="text-gray-500">
                                Reps: {exerciseInfo.repetitions}
                            </p>
                            <p className="text-gray-500">
                                Sets: {exerciseInfo.sets}
                            </p>
                        </div>
                    </div>
                </div>
                <button
                    onClick={handleDeleteExercise}
                    className="cursor-pointer rounded-md bg-red-700 p-2 text-white transition duration-200 hover:bg-red-800"
                >
                    <Trash size={20} />
                </button>
            </div>
        </div>
    );
};

export default ExerciseInfo;
