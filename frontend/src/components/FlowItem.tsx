import { Flow } from '../types';

type FlowItemProps = {
    flow: Flow;
};

const FlowItem = ({ flow }: FlowItemProps) => {
    return (
        <div className="max-w-52 min-w-52 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5">
            <div className="flex h-64 flex-col justify-between rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-5">
                <div className="h-42 rounded-lg bg-gray-300"></div>
                <div className="flex flex-col gap-0.5">
                    <h3 className="text-sm font-bold text-black">{flow.title}</h3>
                    <p className="text-xs text-gray-500">{flow.level}</p>
                </div>
            </div>
        </div>
    );
};

export default FlowItem;
