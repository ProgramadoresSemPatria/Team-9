import { useForm, SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { registerSchema } from '../schemas/register';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router';
import registerUser from '../services/user/register';

type RegisterForm = z.infer<typeof registerSchema>;

const RegisterPage = () => {
    const [isLoading, setIsLoading] = useState(false);

    const navigate = useNavigate();

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<RegisterForm>({
        resolver: zodResolver(registerSchema),
    });

    const onSubmit: SubmitHandler<RegisterForm> = async (registerParams) => {
        setIsLoading(true);
        try {
            console.log(registerParams);

            const response = await registerUser(registerParams);

            if (response?.status !== 201) {
                throw new Error('Error to register');
            }

            console.log(response.data);
            navigate('/sign-in');
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    };
    return (
        <div className="flex h-screen w-full flex-col items-center justify-center gap-4">
            <p className="bg-gradient-to-r from-red-500 to-purple-500 bg-clip-text text-xl font-bold text-transparent">
                GoFit
            </p>
            <h1 className="text-2xl font-bold">Create an account</h1>
            <p className="text-gray-500">
                Already have an account?{' '}
                <Link to="/sign-in" className="text-black hover:text-blue-600">
                    Sign In
                </Link>
            </p>
            <div className="w-96 overflow-hidden rounded-lg bg-white shadow-md">
                <form onSubmit={handleSubmit(onSubmit)} className="space-y-6 p-6">
                    <div className="space-y-4">
                        <div>
                            <label
                                htmlFor="name"
                                className="block text-sm font-medium text-gray-700"
                            >
                                Name
                            </label>
                            <input
                                id="name"
                                type="text"
                                placeholder="Enter your name"
                                className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                    errors.name
                                        ? 'border-red-500'
                                        : 'border-gray-300'
                                }`}
                                {...register('name')}
                            />
                            {errors.name && (
                                <p className="mt-1 text-sm text-red-600">
                                    {errors.name.message}
                                </p>
                            )}
                        </div>
                        <div>
                            <label
                                htmlFor="email"
                                className="block text-sm font-medium text-gray-700"
                            >
                                Email
                            </label>
                            <input
                                id="email"
                                type="email"
                                placeholder="you@example.com"
                                className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                    errors.email
                                        ? 'border-red-500'
                                        : 'border-gray-300'
                                }`}
                                {...register('email')}
                            />
                            {errors.email && (
                                <p className="mt-1 text-sm text-red-600">
                                    {errors.email.message}
                                </p>
                            )}
                        </div>
                        <div>
                            <label
                                htmlFor="password"
                                className="block text-sm font-medium text-gray-700"
                            >
                                Password
                            </label>
                            <input
                                id="password"
                                type="password"
                                placeholder="Create a password"
                                className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                    errors.password
                                        ? 'border-red-500'
                                        : 'border-gray-300'
                                }`}
                                {...register('password')}
                            />
                            {errors.password && (
                                <p className="mt-1 text-sm text-red-600">
                                    {errors.password.message}
                                </p>
                            )}
                        </div>
                        <div>
                            <label
                                htmlFor="ConfirmPassword"
                                className="block text-sm font-medium text-gray-700"
                            >
                                Confirm your password
                            </label>
                            <input
                                id="ConfirmPassword"
                                type="password"
                                placeholder="Confirm password"
                                className={`mt-1 block w-full rounded-md border bg-white px-3 py-2 placeholder-gray-400 shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none ${
                                    errors.password
                                        ? 'border-red-500'
                                        : 'border-gray-300'
                                }`}
                                {...register('confirmPassword')}
                            />
                            {errors.confirmPassword && (
                                <p className="mt-1 text-sm text-red-600">
                                    {errors.confirmPassword.message}
                                </p>
                            )}
                        </div>
                    </div>
                    <div className="pt-2">
                        <button
                            type="submit"
                            disabled={isLoading}
                            className="flex w-full cursor-pointer justify-center rounded-md bg-black px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-gradient-to-r hover:from-red-500 hover:to-purple-500 hover:transition hover:duration-500 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
                        >
                            {isLoading ? 'Creating account...' : 'Sign up'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default RegisterPage;
