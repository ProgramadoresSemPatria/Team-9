import { useForm, SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { createExerciseSchema } from '../schemas/exercise';
import { Exercise } from '../types';

type CreateExerciseForm = z.infer<typeof createExerciseSchema>;

type CreateExerciseFormProps = {
    setExercises: (exercises: Exercise) => void;
};

const CreateExerciseForm = ({ setExercises }: CreateExerciseFormProps) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<CreateExerciseForm>({
        resolver: zodResolver(createExerciseSchema),
    });

    const onSubmit: SubmitHandler<CreateExerciseForm> = async (
        createExerciseParams
    ) => {
        setExercises({ ...createExerciseParams, id: crypto.randomUUID() });
        reset();
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="mt-5 flex flex-col gap-4">
            <div>
                <label htmlFor="title" className="">
                    Title
                </label>
                <input
                    id="title"
                    type="text"
                    placeholder="Ex: Dumbel press"
                    className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                        errors.title ? 'border-red-500' : 'border-gray-300'
                    }`}
                    {...register('title')}
                />
                {errors.title && (
                    <p className="mt-1 text-sm text-red-600">
                        {errors.title.message}
                    </p>
                )}
            </div>
            <div>
                <label htmlFor="muscle" className="">
                    Muscle
                </label>
                <input
                    id="muscle"
                    type="text"
                    placeholder="Ex: Chest"
                    className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                        errors.muscle_group ? 'border-red-500' : 'border-gray-300'
                    }`}
                    {...register('muscle_group')}
                />
                {errors.muscle_group && (
                    <p className="mt-1 text-sm text-red-600">
                        {errors.muscle_group.message}
                    </p>
                )}
            </div>
            <div className="flex w-full items-end justify-between">
                <div>
                    <label htmlFor="repetitions" className="">
                        Repetitions
                    </label>
                    <input
                        id="repetitions"
                        type="number"
                        placeholder="Ex: 12"
                        className={`h-10 w-32 rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                            errors.repetitions ? 'border-red-500' : 'border-gray-300'
                        }`}
                        {...register('repetitions', { valueAsNumber: true })}
                    />
                    {errors.repetitions && (
                        <p className="mt-1 text-sm text-red-600">
                            {errors.repetitions.message}
                        </p>
                    )}
                </div>
                <div>
                    <label htmlFor="sets" className="">
                        Sets
                    </label>
                    <input
                        id="sets"
                        type="text"
                        placeholder="Ex: 3"
                        className={`h-10 w-32 rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                            errors.sets ? 'border-red-500' : 'border-gray-300'
                        }`}
                        {...register('sets', { valueAsNumber: true })}
                    />
                    {errors.sets && (
                        <p className="mt-1 text-sm text-red-600">
                            {errors.sets.message}
                        </p>
                    )}
                </div>
            </div>
            <button
                type="submit"
                className="flex h-10 w-full cursor-pointer items-center justify-center rounded-md bg-black text-white"
            >
                <span className="text-xl">Add</span>
            </button>
        </form>
    );
};

export default CreateExerciseForm;
