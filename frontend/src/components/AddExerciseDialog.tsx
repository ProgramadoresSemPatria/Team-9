import { useState } from 'react';
import CloseDialogBtn from './CloseDialogBtn';
import CreateExerciseForm from './CreateExerciseForm';
import { Exercise } from '../types';

type AddExerciseDialogProps = {
    onClose: () => void;
};

const AddExerciseDialog = ({ onClose }: AddExerciseDialogProps) => {
    const [exercises, setExercises] = useState<Exercise[]>([]);

    return (
        <div className="fixed top-0 flex h-screen w-screen flex-col items-center justify-center backdrop-blur-sm">
            <div className="max-h-[95%] w-80 overflow-auto bg-white p-5">
                <CloseDialogBtn onClick={onClose} />
                <CreateExerciseForm
                    setExercises={(exercise) =>
                        setExercises({ ...exercises, ...exercise })
                    }
                />
            </div>
        </div>
    );
};

export default AddExerciseDialog;
