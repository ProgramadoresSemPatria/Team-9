interface ExerciseInfoProps {
    exerciseInfo: {
        id: string;
        title: string;
        sets: number;
        repetitions: number;
    };
}

const ExerciseInfo = ({ exerciseInfo }: ExerciseInfoProps) => {
    return (
        <div className="rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5">
            <div className="flex items-center gap-3 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-4">
                <div className="h-11 w-11 rounded-md bg-gray-300"></div>
                <div>
                    <h3 className="font-medium text-black">{exerciseInfo.title}</h3>
                    <div className="flex gap-2">
                        <p className="text-gray-500">
                            Reps: {exerciseInfo.repetitions}
                        </p>
                        <p className="text-gray-500">Sets: {exerciseInfo.sets}</p>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ExerciseInfo;
