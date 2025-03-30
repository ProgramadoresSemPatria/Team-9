import trashIcon from '../assets/trash.svg';
import { Exercise } from '../types';

type ExerciseContainerProps = {
    exercise: Exercise;
    onDeleteClick: (id: string) => void;
};

const ExerciseContainer = ({ exercise, onDeleteClick }: ExerciseContainerProps) => {
    return (
        <div key={exercise.id} className="flex w-full items-center justify-center">
            <div className="max-w-xl flex-grow rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5">
                <div className="flex justify-between rounded-md rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-4">
                    <div className="flex flex-col gap-2">
                        <span>
                            {exercise.title} • {exercise.muscle_group}
                        </span>
                        <span className="text-gray-500">
                            Reps: {exercise.repetitions} • Sets: {exercise.sets}
                        </span>
                    </div>

                    <button
                        className="cursor-pointer"
                        onClick={() => onDeleteClick(exercise.id!)}
                    >
                        <img src={trashIcon} alt="trash" />
                    </button>
                </div>
            </div>
        </div>
    );
};

export default ExerciseContainer;
