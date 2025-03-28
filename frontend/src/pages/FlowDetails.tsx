import { useNavigate } from 'react-router';
import plusIcon from '../assets/plus.svg';
import TrainingDayContainer from '../components/TrainingDayContainer';

const daysOfTraining = [
    {
        id: '1',
        title: 'Peito e ombro',
        dayOfWeek: 'Monday',
        exercises: 3,
        duration: 45,
    },
    {
        id: '2',
        title: 'Costas e bíceps',
        dayOfWeek: 'Wednesday',
        exercises: 5,
        duration: 30,
    },
    {
        id: '3',
        title: 'Pernas',
        dayOfWeek: 'Friday',
        exercises: 3,
        duration: 60,
    },
];

const FlowDetailsPage = () => {
    const navigate = useNavigate();

    const handleClick = () => {
        navigate('/add-new-day');
    };

    return (
        <div className="flex w-full flex-col p-7 md:items-center">
            <div className="mb-4 flex w-full flex-col gap-2 md:items-center">
                <h1 className="text-2xl font-bold">Flow name</h1>
                <h3 className="text-xl">Level</h3>
            </div>
            <div className="flex flex-col items-center gap-4">
                {daysOfTraining.length > 0 ? (
                    daysOfTraining.map((trainingDay) => (
                        <TrainingDayContainer
                            key={trainingDay.id}
                            trainingDay={trainingDay}
                        />
                    ))
                ) : (
                    <span>No days of training added</span>
                )}
                <button
                    onClick={handleClick}
                    className="w-full cursor-pointer rounded-md border border-transparent bg-black px-2 py-2 text-xl text-white shadow-sm transition-colors duration-200 hover:bg-gray-800 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:w-64"
                >
                    <div className="relative flex w-full items-center justify-center">
                        <img
                            src={plusIcon}
                            alt="plus"
                            className="absolute left-0 h-6 w-6"
                        />
                        <span className="text-lg">Add training day</span>
                    </div>
                </button>
            </div>
        </div>
    );
};

export default FlowDetailsPage;
