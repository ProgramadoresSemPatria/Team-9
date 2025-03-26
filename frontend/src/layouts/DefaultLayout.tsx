import { Header } from '../components/Header';
import { Outlet } from 'react-router-dom';

export function DefaultLayout() {
    return (
        <div className="flex items-center justify-center">
            <Header />
            <Outlet />
        </div>
    );
}
