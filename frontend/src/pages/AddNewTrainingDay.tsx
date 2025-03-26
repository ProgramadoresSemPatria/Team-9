import CreateNewTrainingDayForm from '../components/CreateNewTrainingDayForm';

const AddNewTrainingDayPage = () => {
    return (
        <div className="container flex flex-col px-7 md:items-center">
            <h1 className="text-2xl font-bold">New Training day</h1>
            <CreateNewTrainingDayForm />
        </div>
    );
};

export default AddNewTrainingDayPage;
