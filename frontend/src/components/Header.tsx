import menuSvg from '../assets/menu.svg';
import userSvg from '../assets/user.svg';
import arrowLeft from '../assets/arrow-left.svg';

export const Header = () => {
    return (
        <header className="flex items-center justify-between p-8">
            <button>
                <img src={menuSvg} alt="menu" />
            </button>
            <p className="text-lg font-bold">GoFit</p>
            <button>
                <img src={userSvg} alt="user" />
            </button>
        </header>
    );
};
