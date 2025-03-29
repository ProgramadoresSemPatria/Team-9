import { Flow } from '../types';

type FlowItemProps = {
    flow: Flow;
};

const FlowItem = ({ flow }: FlowItemProps) => {
    return (
        <div className="flex h-64 max-w-52 min-w-52 flex-col justify-between rounded-lg bg-white p-5">
            <div className="h-42 rounded-lg bg-gray-300"></div>
            <div className="flex flex-col gap-0.5">
                <h3 className="text-sm font-bold text-black">{flow.title}</h3>
                <p className="text-xs text-gray-500">{flow.level}</p>
            </div>
        </div>
    );
};

export default FlowItem;
