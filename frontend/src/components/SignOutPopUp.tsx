import { SignOut } from 'phosphor-react';

type SignOutPopUpProps = {
    handleSignOut: () => void;
};

const SignOutPopUp = ({ handleSignOut }: SignOutPopUpProps) => {
    return (
        <div className="absolute top-16 right-4 z-20 w-48 rounded-lg bg-white p-4 shadow-lg">
            <button
                onClick={handleSignOut}
                className="flex w-full cursor-pointer items-center justify-center gap-3 bg-red-500 text-white hover:bg-red-600"
            >
                <SignOut /> <span>Sign Out</span>
            </button>
        </div>
    );
};

export default SignOutPopUp;
