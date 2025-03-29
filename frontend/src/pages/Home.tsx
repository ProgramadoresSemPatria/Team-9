import searchIcon from '../assets/Search.svg';
import FlowItem from '../components/FlowItem';
import TrainingDayContainer from '../components/TrainingDayContainer';
import { Flow, TrainingDay } from '../types';

const trainingDay: TrainingDay = {
    title: 'Chest',
    dayOfWeek: 'Monday',
    exercises: 5, //Exercise[]
    duration: 45,
};

const flows: Flow[] = [
    {
        id: '1',
        title: 'Aesthetic project',
        level: 'Beginner',
    },
    {
        id: '2',
        title: 'Back focus',
        level: 'Advanced',
    },
    {
        id: '3',
        title: 'Triceps focus',
        level: 'beginner',
    },
    {
        id: '3',
        title: 'Triceps focus',
        level: 'beginner',
    },
    {
        id: '3',
        title: 'Triceps focus',
        level: 'beginner',
    },
];

const HomePage = () => {
    return (
        <>
            <div className="flex w-full flex-col items-center">
                <div className="flex items-center gap-2 p-8">
                    <input
                        type="text"
                        placeholder="Search for flows..."
                        className="h-10 w-72 rounded-md border py-2.5 pl-3"
                    />
                    <button className="cursor-pointer rounded-md border-transparent bg-black px-3 py-3 shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none">
                        <img src={searchIcon} alt="search" />
                    </button>
                </div>
                <div className="h-64 w-full bg-[#808080]"></div>
                <div className="flex h-64 w-full flex-col gap-2.5 p-8">
                    <span className="text-xl font-bold">Today</span>
                    <TrainingDayContainer trainingDay={trainingDay} />
                </div>
                <div className="flex w-full flex-col gap-2.5 p-8 pb-7">
                    <span className="text-xl font-bold">Your flows</span>
                    <div className="flex w-full gap-4 overflow-x-auto">
                        {flows.length > 0 ? (
                            flows.map((flow) => (
                                <FlowItem key={flow.id} flow={flow} />
                            ))
                        ) : (
                            <span>Add a flow</span>
                        )}
                    </div>
                </div>
            </div>
        </>
    );
};

export default HomePage;
