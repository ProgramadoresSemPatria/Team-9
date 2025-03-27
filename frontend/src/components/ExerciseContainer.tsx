import trashIcon from '../assets/trash.svg';
import { Exercise } from '../types';

type ExerciseContainerProps = {
    exercise: Exercise;
    onDeleteClick: (id: string) => void;
};

const ExerciseContainer = ({ exercise, onDeleteClick }: ExerciseContainerProps) => {
    return (
        <div key={exercise.id} className="flex w-full items-center justify-between">
            <div className="flex w-2/3 flex-col gap-2 rounded-md border p-2">
                <span>
                    {exercise.title} • {exercise.muscle}
                </span>
                Reps: {exercise.repetitions} • Sets: {exercise.sets}
            </div>
            <button
                className="cursor-pointer"
                onClick={() => onDeleteClick(exercise.id!)}
            >
                <img src={trashIcon} alt="trash" />
            </button>
        </div>
    );
};

export default ExerciseContainer;
