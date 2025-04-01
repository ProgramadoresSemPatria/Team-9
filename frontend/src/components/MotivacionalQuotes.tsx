import { useEffect, useState } from 'react';

const quotes = [
    'Believe in yourself and all that you are.',
    'Success is the sum of small efforts, repeated day in and day out.',
    'The only way to do great work is to love what you do.',
    'The future belongs to those who believe in the beauty of their dreams.',
    "Don't stop until you're proud.",
    'The only limit to our realization of tomorrow is our doubts of today.',
];

const MotivacionalQuotes = () => {
    const [quote, setQuote] = useState<string>();

    useEffect(() => {
        const quote = getRandomQuote();
        setQuote(quote);
    }, []);

    const getRandomQuote = () => {
        const randomIndex = Math.floor(Math.random() * quotes.length);
        return quotes[randomIndex];
    };

    return (
        <div className="flex h-64 w-full flex-col items-center justify-center bg-[#808080]">
            <div className="flex h-full w-1/2 items-center justify-center">
                <p className="text-center text-2xl font-bold text-black md:text-3xl">
                    {quote}
                </p>
            </div>
        </div>
    );
};

export default MotivacionalQuotes;
