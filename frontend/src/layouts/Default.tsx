import { Outlet, useNavigate } from 'react-router';
import { Header } from '../components/Header';
import Cookies from 'js-cookie';
import { useEffect } from 'react';

const DefaultLayout = () => {
    const navigate = useNavigate();

    const token = Cookies.get('auth_token');

    useEffect(() => {
        if (!token) {
            navigate('/sign-in');
        }
    }, [token, navigate]);

    return (
        <>
            <Header />
            <div className="min-h-screen bg-[#F4F4F4]">
                <Outlet />
            </div>
        </>
    );
};

export default DefaultLayout;
