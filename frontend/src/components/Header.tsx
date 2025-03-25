import menuSvg from '../assets/menu.svg'
import userSvg from '../assets/user.svg'
import arrowLeft from '../assets/arrow-left.svg'

export function Header() {
    return (
        <header className="flex items-center justify-between p-8">
            <button>
                <img src={menuSvg} alt="" />
            </button>
            <p className="text-lg font-bold">GoFit</p>
            <button>
                <img src={userSvg} alt="" />
            </button>
        </header>
    );
}