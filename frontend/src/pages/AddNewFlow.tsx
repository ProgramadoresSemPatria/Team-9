import { z } from 'zod';
import { addNewFlowSchema } from '../schemas/addNewFlow';
import { SubmitHandler, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useState } from 'react';
import { useNavigate } from 'react-router';

type AddNewFlowForm = z.infer<typeof addNewFlowSchema>;

const AddNewFlow = () => {
    const [isLoading, setIsLoading] = useState(false);

    const navigate = useNavigate();

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<AddNewFlowForm>({
        resolver: zodResolver(addNewFlowSchema),
    });

    const onSubmit: SubmitHandler<AddNewFlowForm> = async (registerParams) => {
        setIsLoading(true);
        try {
            console.log(registerParams);
            navigate('/flow-details'); //se a requisição der certo
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <section className="flex min-h-screen w-full flex-col items-center gap-5 px-8">
            <div>
                <h1 className="mt-8 mb-6 w-full text-left text-lg font-bold">
                    New Flow
                </h1>
                <form
                    onSubmit={handleSubmit(onSubmit)}
                    className="flex flex-col gap-4"
                >
                    <div className="flex flex-col gap-1.5">
                        <label
                            htmlFor="newFlowTitle"
                            className="text-sm font-medium"
                        >
                            Title
                        </label>
                        <input
                            type="text"
                            id="newFlowTitle"
                            className={`mt-1 block rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                errors.title ? 'border-red-500' : 'border-gray-300'
                            }`}
                            placeholder="Enter a title for your flow"
                            {...register('title')}
                        />
                        {errors.title && (
                            <p className="mt-1 text-sm text-red-600">
                                {errors.title.message}
                            </p>
                        )}
                    </div>
                    <div className="flex flex-col gap-1.5">
                        <label
                            htmlFor="workoutLevel"
                            className="text-sm font-medium"
                        >
                            Level
                        </label>
                        <select
                            id="workoutLevel"
                            className={`mt-1 block rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                errors.workoutLevel
                                    ? 'border-red-500'
                                    : 'border-gray-300'
                            }`}
                            {...register('workoutLevel')}
                        >
                            <option value="beginner">Beginner</option>
                            <option value="intermediate">Intermediate</option>
                            <option value="advanced">Advanced</option>
                        </select>
                        {errors.workoutLevel && (
                            <p className="mt-1 text-sm text-red-600">
                                {errors.workoutLevel.message}
                            </p>
                        )}
                    </div>
                    <button
                        type="submit"
                        disabled={isLoading}
                        className="mt-2 cursor-pointer rounded-md bg-black px-4 py-2 font-medium text-white"
                    >
                        Create
                    </button>
                </form>
            </div>
        </section>
    );
};

export default AddNewFlow;
