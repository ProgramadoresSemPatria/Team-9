import { Route, Routes } from 'react-router';
import HomePage from './pages/Home';
import RegisterPage from './pages/Register';
import SignInPage from './pages/SignIn';
import HeaderLayout from './layouts/Header';

function App() {
    return (
        <>
            <Routes>
                <Route element={<HeaderLayout />}>
                    <Route path="/" element={<HomePage />} />
                </Route>
                <Route path="register" element={<RegisterPage />} />
                <Route path="sign-in" element={<SignInPage />} />
            </Routes>
        </>
    );
}

export default App;
