import { X } from 'phosphor-react';
import { Link } from 'react-router';
type NavbarModalProps = {
    isNavbarModalopen: boolean;
    setIsNavbarModalopen: () => void;
};

const NavbarModal = ({
    isNavbarModalopen,
    setIsNavbarModalopen,
}: NavbarModalProps) => {
    return (
        <div
            className={`fixed top-0 left-0 z-30 h-full w-64 transform bg-white shadow-lg transition-transform duration-300 ease-in-out ${
                isNavbarModalopen ? 'translate-x-0' : '-translate-x-full'
            }`}
        >
            <div className="flex items-center justify-between border-b p-4">
                <h2 className="text-lg font-semibold">Navigation</h2>
                <button
                    onClick={() => setIsNavbarModalopen()}
                    className="cursor-pointer rounded-lg p-1 hover:bg-gray-200"
                >
                    <X size={24} />
                </button>
            </div>
            <nav className="flex flex-col space-y-3 p-4">
                <Link to="/" className="text-gray-700 hover:text-blue-600">
                    Home
                </Link>
                <Link
                    to="/add-new-flow"
                    className="text-gray-700 hover:text-blue-600"
                >
                    Add new flow
                </Link>
            </nav>
        </div>
    );
};

export default NavbarModal;
