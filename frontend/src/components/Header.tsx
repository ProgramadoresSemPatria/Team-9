import { useState } from 'react';
import menuSvg from '../assets/menu.svg';
import userSvg from '../assets/user.svg';
import { useNavigate } from 'react-router';
import Cookies from 'js-cookie';
import SignOutPopUp from './SignOutPopUp';
import NavbarModal from './NavbarModal';

export const Header = () => {
    const [isSignOutModalopen, setIsSignOutModalopen] = useState(false);
    const [isNavbarModalopen, setIsNavbarModalopen] = useState(false);

    const navigate = useNavigate();

    const handleSignOut = () => {
        navigate('/sign-in');
        Cookies.remove('auth_token');
    };

    return (
        <div>
            <header className="relative z-10 flex items-center justify-between bg-gradient-to-r from-red-500 to-purple-500 bg-clip-text p-8 text-transparent">
                <div className="mx-auto flex w-full max-w-7xl items-center justify-between">
                    <button
                        onClick={() => setIsNavbarModalopen(true)}
                        className="cursor-pointer"
                    >
                        <img src={menuSvg} alt="menu" />
                    </button>
                    <div className="flex items-center justify-center">
                        <p className="text-xl font-bold">GoFit</p>
                    </div>
                    <button
                        onClick={() => setIsSignOutModalopen(!isSignOutModalopen)}
                        className="cursor-pointer"
                    >
                        <img src={userSvg} alt="user" />
                    </button>
                </div>
                {isSignOutModalopen && (
                    <SignOutPopUp handleSignOut={handleSignOut} />
                )}
            </header>

            <NavbarModal
                isNavbarModalopen={isNavbarModalopen}
                setIsNavbarModalopen={() => setIsNavbarModalopen(false)}
            />

            {isNavbarModalopen && (
                <div
                    className="fixed inset-0 z-20 bg-black/30 backdrop-blur-md"
                    onClick={() => setIsNavbarModalopen(false)}
                ></div>
            )}
        </div>
    );
};
