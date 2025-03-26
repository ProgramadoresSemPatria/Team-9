const CloseDialogBtn = ({ onClick }: { onClick: () => void }) => {
    return (
        <button
            className="text-brand-secondary hover:bg-brand-neutral hover:text-brand-accent rounded-sm border-2 px-2 transition-colors duration-200"
            onClick={onClick}
        >
            X
        </button>
    );
};

export default CloseDialogBtn;
