import { Flow } from '../types';
import { Barbell } from 'phosphor-react';

type FlowItemProps = {
    flow: Flow;
};

const FlowItem = ({ flow }: FlowItemProps) => {
    return (
        <div className="w-56 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5 md:w-64">
            <div className="flex items-center gap-4 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-5">
                <div className="h-11 w-11 rounded-md bg-zinc-200 p-3">
                    <Barbell size={20} weight="fill" className="text-purple-500" />
                </div>
                <div className="flex flex-col gap-0.5">
                    <h3 className="text-sm font-bold text-black">{flow.title}</h3>
                    <p className="text-xs text-gray-500">{flow.level}</p>
                </div>
            </div>
        </div>
    );
};

export default FlowItem;
