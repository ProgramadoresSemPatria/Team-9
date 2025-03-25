const AddNewFlow = () => {
    return (
        <section className="flex min-h-screen w-full flex-col items-center gap-5">
            <h1 className="text-lg font-bold">New Flow</h1>
            <div>
                <form className="flex flex-col gap-4">
                    <div className="flex flex-col gap-1.5">
                        <label
                            htmlFor="newFlowTitle"
                            className="text-sm font-medium"
                        >
                            Title
                        </label>
                        <input
                            type="text"
                            id="newFlowTitle"
                            className="rounded-md border-1 border-gray-200 px-3 py-2.5"
                            placeholder="Enter a title for your flow"
                        />
                    </div>
                    <div className="flex flex-col gap-1.5">
                        <label
                            htmlFor="workoutLevel"
                            className="text-sm font-medium"
                        >
                            Level
                        </label>
                        <select
                            name=""
                            id="workoutLevel"
                            className="rounded-md border-1 border-gray-200 px-3 py-2.5"
                        >
                            <option value="beginner">Beginner</option>
                            <option value="intermediate">Intermediate</option>
                            <option value="advanced">Advanced</option>
                        </select>
                    </div>
                    <div className="flex flex-col gap-1.5">
                        <label htmlFor="coverImg" className="text-sm font-medium">
                            Cover
                        </label>
                        <input
                            type="file"
                            name=""
                            id=""
                            className="rounded-md border-1 border-gray-200 px-3 py-2.5"
                        />
                    </div>
                </form>
            </div>
            <button className="w-[337px] rounded-md bg-black px-4 py-2 font-medium text-white">
                Criar
            </button>
        </section>
    );
};

export default AddNewFlow;
