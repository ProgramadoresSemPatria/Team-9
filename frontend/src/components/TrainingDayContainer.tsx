import { TrainingDay } from '../types';
import calendarIcon from '../assets/Calendar.svg';
import fireIcon from '../assets/Vector.svg';
import clockIcon from '../assets/Group.svg';

type TrainingDayContainerProps = {
    trainingDay: TrainingDay;
};

const TrainingDayContainer = ({ trainingDay }: TrainingDayContainerProps) => {
    return (
        <div
            key={trainingDay.id}
            className="flex h-44 w-full flex-col justify-end gap-1 rounded-xl border p-5 shadow-sm md:w-96"
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
    );
};

export default TrainingDayContainer;
