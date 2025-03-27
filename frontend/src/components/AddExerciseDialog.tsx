import CloseDialogBtn from './CloseDialogBtn';
import CreateExerciseForm from './CreateExerciseForm';
import { Exercise } from '../types';

type AddExerciseDialogProps = {
    onClose: () => void;
    setExercises: (exercise: Exercise) => void;
};

const AddExerciseDialog = ({ onClose, setExercises }: AddExerciseDialogProps) => {
    return (
        <div className="fixed top-0 flex h-screen w-screen flex-col items-center justify-center backdrop-blur-sm">
            <div className="max-h-[95%] w-80 overflow-auto rounded-md bg-white p-5">
                <CloseDialogBtn onClick={onClose} />
                <CreateExerciseForm setExercises={setExercises} />
            </div>
        </div>
    );
};

export default AddExerciseDialog;
