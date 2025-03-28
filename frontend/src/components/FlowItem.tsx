const FlowItem = () => {
    return (
        <div className="flex h-64 w-52 flex-col justify-between rounded-lg bg-zinc-100 p-5">
            <div className="h-42 rounded-lg bg-gray-300"></div>
            <div className="flex flex-col gap-0.5">
                <h3 className="text-sm font-bold text-black">Focus on chest</h3>
                <p className="text-xs text-gray-500">Intermediate</p>
            </div>
        </div>
    );
};

export default FlowItem;
