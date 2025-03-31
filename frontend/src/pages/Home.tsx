import { useEffect, useState } from 'react';
import searchIcon from '../assets/Search.svg';
import FlowItem from '../components/FlowItem';
import TrainingDayContainer from '../components/TrainingDayContainer';
import { Flow, TrainingDay } from '../types';
import getFlowsByUser from '../services/flows/getAll';
import Cookies from 'js-cookie';
import { Link, useNavigate } from 'react-router';
import MotivacionalQuotes from '../components/MotivacionalQuotes';

const trainingDay: TrainingDay = {
    title: 'Chest',
    dayOfWeek: 'Monday',
    exercises: 5, //Exercise[]
    duration: 45,
};

const HomePage = () => {
    const [inputValue, setInputValue] = useState('');
    const [flowsList, setFlowsList] = useState<Flow[]>([]);
    const [flowsFinds, setFlowsFinds] = useState<Flow[]>();

    const navigate = useNavigate();

    useEffect(() => {
        const getFlows = async () => {
            const token = Cookies.get('auth_token');

            if (!token) throw new Error('JWT token invalid');

            const response = await getFlowsByUser(token);

            if (response?.status !== 200) {
                throw new Error('Error to get flows');
            }

            setFlowsList(response.data);
        };
        getFlows();
    }, []);

    const handleSearchFlow = () => {
        if (!inputValue.trim()) {
            return;
        }
        const flows = flowsList.filter((flow) =>
            flow.title.toLowerCase().includes(inputValue.toLowerCase())
        );
        setFlowsFinds(flows);
        setInputValue('');
    };

    return (
        <>
            <div className="flex w-full flex-col items-center">
                <div className="flex items-center justify-center gap-2 p-8 md:w-full md:max-w-6xl">
                    <input
                        type="text"
                        placeholder="Search for flows..."
                        value={inputValue}
                        onChange={(e) => setInputValue(e.target.value)}
                        className="h-10 w-72 flex-grow rounded-md border-2 border-gray-200 bg-white py-3 pl-3"
                    />
                    <button
                        onClick={handleSearchFlow}
                        className="flex cursor-pointer items-center justify-center gap-2 rounded-md border-transparent bg-black px-3 py-3 shadow-sm transition-colors duration-200 hover:bg-gray-800 hover:bg-gradient-to-r hover:from-red-500 hover:to-purple-500 hover:transition hover:duration-500 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none md:py-2"
                    >
                        <img src={searchIcon} alt="search" />
                        <span className="hidden text-white md:inline-block">
                            Search
                        </span>
                    </button>
                </div>
                <div className="mb-3 flex w-full gap-4 overflow-x-auto px-8">
                    {flowsFinds &&
                        flowsFinds.map((flow) => (
                            <Link key={flow.id} to={`/flow-details/${flow.id}`}>
                                <FlowItem key={flow.id} flow={flow} />
                            </Link>
                        ))}
                </div>
                <MotivacionalQuotes />
                <div className="flex h-64 w-full max-w-7xl flex-col gap-2.5 p-8">
                    <span className="text-xl font-bold">Today</span>
                    <TrainingDayContainer trainingDay={trainingDay} />
                </div>
                <div className="flex w-full max-w-7xl flex-col gap-2.5 p-8 pb-7">
                    <span className="text-xl font-bold">Your flows</span>
                    <div className="flex w-full gap-4 overflow-x-auto">
                        {flowsList.length > 0 ? (
                            flowsList.map((flow) => (
                                <Link key={flow.id} to={`/flow-details/${flow.id}`}>
                                    <FlowItem key={flow.id} flow={flow} />
                                </Link>
                            ))
                        ) : (
                            <button
                                className="cursor-pointer rounded-md border-transparent bg-black px-3 py-3 text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none"
                                onClick={() => navigate('/add-new-flow')}
                            >
                                Add a flow
                            </button>
                        )}
                    </div>
                </div>
            </div>
        </>
    );
};

export default HomePage;
