import { useState } from 'react';
import CreateNewTrainingDayForm from '../components/CreateNewTrainingDayForm';
import { createPortal } from 'react-dom';
import AddExerciseDialog from '../components/AddExerciseDialog';
import { Exercise } from '../types';
import ExerciseContainer from '../components/ExerciseContainer';

const AddNewTrainingDayPage = () => {
    const [addExerciseDialogIsOpen, setAddExerciseDialogIsOpen] = useState(false);
    const [exercises, setExercises] = useState<Exercise[]>([]);

    const handleSetExercises = (newExercise: Exercise) => {
        setExercises((prevExercises) => [...prevExercises, newExercise]);
    };

    const handleDeleteExercise = (id: string) => {
        setExercises(exercises.filter((exercise) => exercise.id !== id));
    };

    return (
        <div className="container mx-auto flex w-full flex-col gap-2 px-7 py-5 md:items-center">
            <h1 className="text-left text-2xl font-bold">New Training day</h1>
            <CreateNewTrainingDayForm
                exercises={exercises}
                setOpenAddExerciseDialog={() => setAddExerciseDialogIsOpen(true)}
            />
            <div className="my-5 flex w-full flex-col items-center gap-3 border-y-2 border-gray-300 py-5 md:w-1/2">
                <h2 className="text-xl font-bold">Exercises</h2>
                <div className="flex h-56 w-full flex-col items-start gap-2 overflow-auto">
                    {exercises.length > 0 ? (
                        exercises.map((exercise) => (
                            <ExerciseContainer
                                key={exercise.id}
                                exercise={exercise}
                                onDeleteClick={handleDeleteExercise}
                            />
                        ))
                    ) : (
                        <div className="flex w-full items-center justify-center">
                            <span>No exercises added</span>
                        </div>
                    )}
                </div>
            </div>
            {addExerciseDialogIsOpen &&
                createPortal(
                    <AddExerciseDialog
                        onClose={() => setAddExerciseDialogIsOpen(false)}
                        setExercises={handleSetExercises}
                    />,
                    document.body
                )}
        </div>
    );
};

export default AddNewTrainingDayPage;
