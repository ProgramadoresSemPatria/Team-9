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
        <div className="flex w-full flex-col items-center p-7">
            <div className="flex flex-col items-center gap-4">
                <div className="mb-4 flex w-full flex-col gap-2">
                    <h1 className="text-2xl font-bold">Flow name</h1>
                    <h3 className="text-xl text-gray-500">Level</h3>
                </div>
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
                    className="flex w-full cursor-pointer items-center justify-center rounded-md bg-white p-3 text-xl text-black shadow-sm transition-colors duration-200 hover:bg-gradient-to-r hover:from-red-500 hover:to-purple-500 hover:text-white hover:transition hover:duration-500 focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
                >
                    <div className="flex items-center justify-center rounded-md bg-black p-3">
                        <img src={plusIcon} alt="plus" className="h-6 w-6" />
                    </div>
                    <span className="flex-grow text-lg font-medium">
                        Add training day
                    </span>
                </button>
            </div>
        </div>
    );
};

export default FlowDetailsPage;
