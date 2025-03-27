import { useState } from 'react';
import CreateNewTrainingDayForm from '../components/CreateNewTrainingDayForm';
import { createPortal } from 'react-dom';
import AddExerciseDialog from '../components/AddExerciseDialog';

const AddNewTrainingDayPage = () => {
    const [addExerciseDialogIsOpen, setAddExerciseDialogIsOpen] = useState(false);
    return (
        <div className="container flex flex-col gap-2 px-7 md:items-center">
            <h1 className="text-2xl font-bold">New Training day</h1>
            <CreateNewTrainingDayForm />
            <button
                className="cursor-pointer rounded-md border border-transparent bg-black px-4 py-2 text-xl text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:w-48"
                onClick={() => setAddExerciseDialogIsOpen(true)}
            >
                Add exercise
            </button>
            <div className="flex w-full flex-col items-center border-y-2 border-gray-300 py-5">
                <h2 className="text-xl font-bold">Exercises</h2>
                <span>No exercises added</span>
            </div>
            {addExerciseDialogIsOpen &&
                createPortal(
                    <AddExerciseDialog
                        onClose={() => setAddExerciseDialogIsOpen(false)}
                    />,
                    document.body
                )}
        </div>
    );
};

export default AddNewTrainingDayPage;
