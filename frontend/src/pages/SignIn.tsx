import { useForm, SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router';
import { signInSchema } from '../schemas/signIn';
import login from '../services/user/login';
import Cookies from 'js-cookie';

type SignInForm = z.infer<typeof signInSchema>;

const SignInPage = () => {
    const [isLoading, setIsLoading] = useState(false);

    const navigate = useNavigate();

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<SignInForm>({
        resolver: zodResolver(signInSchema),
    });

    const onSubmit: SubmitHandler<SignInForm> = async (signInParams) => {
        setIsLoading(true);
        try {
            const response = await login(signInParams);

            if (response?.status !== 200) {
                throw new Error('Error to log in');
            }
            console.log(response.data);
            const token = response.data.token;

            Cookies.set('auth_token', token, { expires: 7 });
            navigate('/');
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    };
    return (
        <div className="flex h-screen w-full flex-col items-center justify-center gap-4">
            <h1 className="text-4xl font-bold">Sign in your account</h1>

            <div className="w-96 overflow-hidden rounded-lg bg-white shadow-md">
                <form onSubmit={handleSubmit(onSubmit)} className="space-y-6 p-6">
                    <div className="space-y-4">
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
                                placeholder="Your password"
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
                    </div>
                    <div className="pt-2">
                        <button
                            type="submit"
                            disabled={isLoading}
                            className="flex w-full cursor-pointer justify-center rounded-md border border-transparent bg-black px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
                        >
                            {isLoading ? 'Creating account...' : 'Sign in'}
                        </button>
                    </div>
                    <p>
                        Don&apos;t have an account?{' '}
                        <Link to="/register" className="hover:text-blue-600">
                            Sign up
                        </Link>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default SignInPage;
