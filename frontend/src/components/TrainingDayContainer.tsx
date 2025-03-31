import { TrainingDay } from '../types';
import calendarIcon from '../assets/Calendar.svg';
import fireIcon from '../assets/Vector.svg';
import clockIcon from '../assets/Group.svg';

type TrainingDayContainerProps = {
    trainingDay: TrainingDay;
};

const TrainingDayContainer = ({ trainingDay }: TrainingDayContainerProps) => {
    return (
        <div className="w-80 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5 transition-all hover:scale-105 hover:rotate-1 hover:duration-500 md:w-md">
            <div
                key={trainingDay.id}
                className="flex h-44 flex-col justify-end gap-1 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-5"
            >
                <div className="flex gap-2">
                    <img src={calendarIcon} alt="calendar" />
                    <span>{trainingDay.dayOfWeek}</span>
                </div>
                <span className="font-bold">{trainingDay.title}</span>
                <div className="flex items-center gap-2">
                    <div className="flex items-center gap-2">
                        <img src={fireIcon} alt="fire" />
                        <span>{trainingDay.exercises} exercises</span>
                    </div>
                    <div className="flex items-center gap-2">
                        <img src={clockIcon} alt="clock" />
                        <span>{trainingDay.duration} min</span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default TrainingDayContainer;
