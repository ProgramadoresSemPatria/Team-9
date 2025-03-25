import { Outlet } from 'react-router';
import { Header } from '../components/Header';

const HeaderLayout = () => {
    return (
        <>
            <Header />
            <Outlet />
        </>
    );
};

export default HeaderLayout;
