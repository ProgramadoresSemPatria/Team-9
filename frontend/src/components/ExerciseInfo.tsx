const ExerciseInfo = () => {
    return (
        <div className="flex items-center gap-3 rounded-lg bg-zinc-100 p-4">
            <div className="h-11 w-11 rounded-md bg-gray-300"></div>
            <div>
                <h3 className="font-medium text-black">Exercise Title</h3>
                <div className="flex gap-2">
                    <p className="text-gray-500">Reps: </p>
                    <p className="text-gray-500">Sets: </p>
                </div>
            </div>
        </div>
    );
};

export default ExerciseInfo;
