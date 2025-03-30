import menuSvg from '../assets/menu.svg';
import userSvg from '../assets/user.svg';

export const Header = () => {
    return (
        <header className="flex items-center justify-between bg-gradient-to-r from-red-500 to-purple-500 bg-clip-text p-8 text-transparent">
            <div className="mx-auto flex w-full max-w-7xl items-center justify-between">
                <button>
                    <img src={menuSvg} alt="menu" />
                </button>
                <div className="flex items-center justify-center">
                    <p className="text-lg font-bold">GoFit</p>
                </div>
                <button>
                    <img src={userSvg} alt="user" />
                </button>
            </div>
        </header>
    );
};
