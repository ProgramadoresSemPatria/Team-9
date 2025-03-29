import { Outlet } from 'react-router';
import { Header } from '../components/Header';

const HeaderLayout = () => {
    return (
        <>
            <Header />
            <div className="bg-[#F4F4F4]">
                <Outlet />
            </div>
        </>
    );
};

export default HeaderLayout;
