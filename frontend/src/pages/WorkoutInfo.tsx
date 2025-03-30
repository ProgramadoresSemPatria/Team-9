import { CheckCircle, Fire, Timer } from 'phosphor-react';
import ExerciseInfo from '../components/ExerciseInfo';
import { useState } from 'react';

const WorkoutInfo = () => {
    const [workout, setWorkout] = useState({
        id: 1,
        name: 'Full Body Workout',
        duration: 45,
        isWorkoutFinished: false,
        exercises: [
            { id: 1, name: 'Push-up', sets: 3, reps: 15 },
            { id: 2, name: 'Squat', sets: 3, reps: 20 },
            { id: 3, name: 'Deadlift', sets: 3, reps: 10 },
        ],
    });

    const handleFinishWorkout = () => {
        setWorkout((prev) => ({ ...prev, isWorkoutFinished: true }));
    };

    return (
        <div className="mx-auto flex min-h-screen max-w-4xl flex-col gap-5 p-8">
            <div className="flex flex-col gap-1.5">
                <h2 className="text-lg font-bold">{workout.name}</h2>
                <div className="flex gap-3">
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Fire weight="fill" className="text-gray-400" />
                            <p>
                                <span>{workout.exercises.length}</span> exercícios
                            </p>
                        </span>
                    </div>
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Timer weight="fill" className="text-gray-400" />
                            <p>
                                <span>{workout.duration}</span> minutos
                            </p>
                        </span>
                    </div>
                </div>
            </div>
            <div className="flex flex-col gap-5">
                {workout.exercises.map((exercise) => (
                    <ExerciseInfo key={exercise.name} exerciseInfo={exercise} />
                ))}
            </div>
            <button
                className="flex w-full cursor-pointer items-center justify-center gap-1 rounded-md bg-black px-4 py-2 font-medium text-white"
                onClick={handleFinishWorkout}
            >
                Mark as finished
                <CheckCircle weight="fill" size={18} className="text-white" />
            </button>
        </div>
    );
};

export default WorkoutInfo;
