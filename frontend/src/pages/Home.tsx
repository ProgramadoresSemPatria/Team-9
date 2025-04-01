import { useEffect, useState } from 'react';
import searchIcon from '../assets/Search.svg';
import FlowItem from '../components/FlowItem';
import { Flow } from '../types';
import getFlowsByUser from '../services/flows/getAll';
import Cookies from 'js-cookie';
import { Link, useNavigate } from 'react-router';
import MotivacionalQuotes from '../components/MotivacionalQuotes';
import { Eraser } from 'phosphor-react';

const HomePage = () => {
    const [inputValue, setInputValue] = useState('');
    const [flowsList, setFlowsList] = useState<Flow[]>([]);
    const [flowsFinds, setFlowsFinds] = useState<Flow[] | null>(null);

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
        const trimmedInput = inputValue.trim();
        if (!trimmedInput) {
            setFlowsFinds(null);
            return;
        }
        const flows = flowsList.filter((flow) =>
            flow.title.toLowerCase().includes(trimmedInput.toLowerCase())
        );
        setFlowsFinds(flows);
    };

    const handleClearSearch = () => {
        setInputValue('');
        setFlowsFinds(null);
    };

    return (
        <div className="flex w-full flex-col items-center">
            <div className="w-full shadow-md">
                <div className="mx-auto flex items-center justify-center gap-2 p-4 md:w-full md:max-w-6xl">
                    <input
                        type="text"
                        placeholder="Search for flows..."
                        value={inputValue}
                        onChange={(e) => setInputValue(e.target.value)}
                        className="h-10 w-full max-w-md flex-grow rounded-xl border border-gray-300 bg-white px-4 py-2 shadow-sm focus:border-purple-500 focus:ring-2 focus:ring-purple-300 focus:outline-none"
                    />
                    <button
                        onClick={handleSearchFlow}
                        className="flex cursor-pointer items-center justify-center gap-2 rounded-md border-transparent bg-black px-3 py-3 text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 hover:bg-gradient-to-r hover:from-red-500 hover:to-purple-500 hover:transition hover:duration-500 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none md:py-2"
                    >
                        <img src={searchIcon} alt="search" className="h-5 w-5" />
                        <span className="hidden md:inline-block">Search</span>
                    </button>
                    {flowsFinds && (
                        <button
                            onClick={handleClearSearch}
                            className="flex cursor-pointer items-center gap-2.5 rounded-md border px-2 py-1.5 text-sm text-gray-600 shadow-sm hover:text-black"
                        >
                            <Eraser size={32} />
                            <span>Clear</span>
                        </button>
                    )}
                </div>
            </div>

            {flowsFinds && (
                <div className="w-full max-w-7xl px-8">
                    <h2 className="mt-6 mb-3 text-lg font-semibold">
                        Search results for:{' '}
                        <span className="italic">{inputValue}</span>
                    </h2>
                    <div className="mb-6 flex gap-4 overflow-x-auto transition-opacity duration-500 ease-in">
                        {flowsFinds.length > 0 ? (
                            flowsFinds.map((flow) => (
                                <Link key={flow.id} to={`/flow-details/${flow.id}`}>
                                    <FlowItem flow={flow} />
                                </Link>
                            ))
                        ) : (
                            <span className="text-gray-500">No flows found.</span>
                        )}
                    </div>
                </div>
            )}

            <MotivacionalQuotes />

            <div className="mt-3 flex w-full max-w-7xl flex-col gap-3 px-8 pb-10">
                <div className="flex items-center justify-between">
                    <span className="text-xl font-bold">Your flows</span>
                    <button
                        onClick={() => navigate('/add-new-flow')}
                        className="flex cursor-pointer items-center justify-center gap-2 rounded-md border-transparent bg-black px-3 py-3 text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 hover:bg-gradient-to-r hover:from-red-500 hover:to-purple-500 hover:transition hover:duration-500 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none md:py-2"
                    >
                        Add a flow
                    </button>
                </div>

                <div className="flex gap-4 overflow-x-auto transition-opacity duration-500 ease-in">
                    {flowsList.length > 0 ? (
                        flowsList.map((flow) => (
                            <Link key={flow.id} to={`/flow-details/${flow.id}`}>
                                <FlowItem flow={flow} />
                            </Link>
                        ))
                    ) : (
                        <span className="text-gray-500">No flows added yet.</span>
                    )}
                </div>
            </div>
        </div>
    );
};

export default HomePage;
