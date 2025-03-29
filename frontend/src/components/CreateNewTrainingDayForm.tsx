import { useForm, SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { createNewTrainingDaySchema } from '../schemas/trainingDay';
import { Exercise } from '../types';
import { useState } from 'react';
import Cookies from 'js-cookie';
import createTrainingDay from '../services/trainingDay/create';
import { useNavigate, useParams } from 'react-router';

type CreateNewTrainingDayFormType = z.infer<typeof createNewTrainingDaySchema>;

type CreateNewTrainingDayFormProps = {
    exercises: Exercise[];
    setOpenAddExerciseDialog: () => void;
};

const CreateNewTrainingDayForm = ({
    exercises,
    setOpenAddExerciseDialog,
}: CreateNewTrainingDayFormProps) => {
    const [isLoading, setIsLoading] = useState(false);

    const { id } = useParams();

    if (!id) return;

    const navigate = useNavigate();

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<CreateNewTrainingDayFormType>({
        resolver: zodResolver(createNewTrainingDaySchema),
    });

    const onSubmit: SubmitHandler<CreateNewTrainingDayFormType> = async (
        createNewTrainingDayParams
    ) => {
        console.log(createNewTrainingDayParams);
        console.log(exercises);
        setIsLoading(true);
        try {
            const token = Cookies.get('auth_token');

            if (!token) throw new Error('JWT token invalid');

            const response = await createTrainingDay(
                { ...createNewTrainingDayParams, flowId: id },
                token
            );

            if (response?.status !== 201) {
                throw new Error('Error to create flow');
            }

            console.log(response.data);
            navigate('/flow-details');
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    };
    return (
        <form
            onSubmit={handleSubmit(onSubmit)}
            className="mt-5 flex flex-col gap-4 md:w-80"
        >
            <div>
                <label htmlFor="title" className="">
                    Title
                </label>
                <input
                    id="title"
                    type="text"
                    placeholder="Ex: Chest and legs"
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
            <div className="flex w-full items-start justify-between px-3 md:justify-between md:px-0">
                <div className="flex flex-col">
                    <label htmlFor="day" className="">
                        Day
                    </label>
                    <select
                        id="day"
                        className={`mt-1 h-10 w-36 rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                            errors.dayOfWeek ? 'border-red-500' : 'border-gray-300'
                        }`}
                        {...register('dayOfWeek')}
                    >
                        <option value="" disabled>
                            Select a day
                        </option>
                        <option value="Monday">Monday</option>
                        <option value="Tuesday">Tuesday</option>
                        <option value="Wednesday">Wednesday</option>
                        <option value="Thursday">Thursday</option>
                        <option value="Friday">Friday</option>
                        <option value="Saturday">Saturday</option>
                        <option value="Sunday">Sunday</option>
                    </select>
                    {errors.dayOfWeek && (
                        <p className="mt-1 text-sm text-red-600">
                            {errors.dayOfWeek.message}
                        </p>
                    )}
                </div>
                <div className="flex flex-col">
                    <label htmlFor="duration" className="">
                        Duration (minutes)
                    </label>
                    <input
                        id="sets"
                        type="text"
                        placeholder="Ex: 3"
                        className={`mt-1 h-10 w-36 rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                            errors.duration ? 'border-red-500' : 'border-gray-300'
                        }`}
                        {...register('duration', { valueAsNumber: true })}
                    />
                    {errors.duration && (
                        <p className="mt-1 text-sm text-red-600">
                            {errors.duration.message}
                        </p>
                    )}
                </div>
            </div>
            <div className="flex w-full flex-col items-center gap-3">
                <button
                    className="cursor-pointer rounded-md border border-transparent bg-black px-4 py-2 text-xl text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:w-48"
                    onClick={() => setOpenAddExerciseDialog()}
                    disabled={isLoading}
                >
                    Add exercise
                </button>
                <button
                    type="submit"
                    disabled={isLoading}
                    className="flex h-10 w-full cursor-pointer items-center justify-center rounded-md bg-black text-white md:w-64"
                >
                    <span className="text-xl">Save</span>
                </button>
            </div>
        </form>
    );
};

export default CreateNewTrainingDayForm;
