/* eslint-disable prettier/prettier */
import { CheckCircle, Fire, Timer } from 'phosphor-react';

const WorkoutInfo = () => {
    return (
        <div className="flex flex-col gap-5 p-8">
            <div className="flex flex-col gap-1.5">
                <h2 className="text-lg font-bold">Workout Title</h2>
                <div className="flex gap-3">
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Fire weight="fill" className="text-gray-400" />
                            <p>
                                <span>3</span> exercises
                            </p>
                        </span>
                    </div>
                    <div>
                        <span className="flex items-center gap-1 text-gray-400">
                            <Timer weight="fill" className="text-gray-400" />
                            <p>
                                <span>45</span> minutes
                            </p>
                        </span>
                    </div>
                </div>
            </div>
            <div className="flex flex-col gap-5">
                
            </div>
            <button className="flex w-full items-center justify-center gap-1 rounded-md bg-black px-4 py-2 font-medium text-white">
                Mark as finished
                <CheckCircle weight="fill" size={18} className="text-white" />
            </button>
        </div>
    );
};

export default WorkoutInfo;
