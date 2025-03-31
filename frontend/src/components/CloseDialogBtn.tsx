const CloseDialogBtn = ({ onClick }: { onClick: () => void }) => {
    return (
        <button
            className="cursor-pointer rounded-sm border-2 border-none bg-red-600 px-2 text-white transition-colors duration-200 hover:bg-red-800"
            onClick={onClick}
        >
            X
        </button>
    );
};

export default CloseDialogBtn;
