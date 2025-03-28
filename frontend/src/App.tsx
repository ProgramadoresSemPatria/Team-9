import { Route, Routes } from 'react-router';
import HomePage from './pages/Home';
import RegisterPage from './pages/Register';
import HeaderLayout from './layouts/Header';
import SignInPage from './pages/SignIn';
import AddNewFlowPage from './pages/AddNewFlow';
import AddNewTrainingDayPage from './pages/AddNewTrainingDay';
import WorkoutInfo from './pages/WorkoutInfo';
import FlowDetailsPage from './pages/FlowDetails';

function App() {
    return (
        <>
            <Routes>
                <Route element={<HeaderLayout />}>
                    <Route path="/" element={<HomePage />} />
                    <Route path="add-new-flow" element={<AddNewFlowPage />} />
                    <Route path="add-new-day" element={<AddNewTrainingDayPage />} />
                    <Route path="workout-info" element={<WorkoutInfo />} />
                    <Route path="flow-details" element={<FlowDetailsPage />} />
                </Route>
                <Route path="register" element={<RegisterPage />} />
                <Route path="sign-in" element={<SignInPage />} />
            </Routes>
        </>
    );
}

export default App;
