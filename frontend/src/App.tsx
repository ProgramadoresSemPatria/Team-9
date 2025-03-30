import { Route, Routes } from 'react-router';
import HomePage from './pages/Home';
import RegisterPage from './pages/Register';
import SignInPage from './pages/SignIn';
import AddNewFlowPage from './pages/AddNewFlow';
import AddNewTrainingDayPage from './pages/AddNewTrainingDay';
import WorkoutInfo from './pages/WorkoutInfo';
import FlowDetailsPage from './pages/FlowDetails';
import DefaultLayout from './layouts/Default';

function App() {
    return (
        <>
            <Routes>
                <Route element={<DefaultLayout />}>
                    <Route path="/" element={<HomePage />} />
                    <Route path="add-new-flow" element={<AddNewFlowPage />} />
                    <Route
                        path="add-new-day/:id"
                        element={<AddNewTrainingDayPage />}
                    />
                    <Route path="workout-info" element={<WorkoutInfo />} />
                    <Route path="flow-details/:id" element={<FlowDetailsPage />} />
                </Route>
                <Route path="register" element={<RegisterPage />} />
                <Route path="sign-in" element={<SignInPage />} />
            </Routes>
        </>
    );
}

export default App;
