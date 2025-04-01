import { Flow } from '../types';
import { Barbell } from 'phosphor-react';
import { motion } from 'framer-motion';

type FlowItemProps = {
    flow: Flow;
};

const FlowItem = ({ flow }: FlowItemProps) => {
    return (
        <motion.div
            layout
            initial={{ opacity: 0, y: 30 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: 20 }}
            transition={{
                duration: 0.6,
                ease: [0.34, 1.56, 0.64, 1],
            }}
            whileHover={{
                scale: 1.04,
                background:
                    'linear-gradient(135deg, rgba(239, 68, 68, 1), rgba(168, 85, 247, 1))',
                boxShadow:
                    '0 0 12px rgba(168, 85, 247, 0.4), 0 12px 28px rgba(0, 0, 0, 0.15)',
                transition: { duration: 0.2 },
            }}
            whileTap={{ scale: 0.98 }}
            className="w-56 cursor-pointer rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-gradient-to-br from-red-500 to-purple-500 p-0.5 md:w-64"
        >
            <div className="flex items-center gap-4 rounded-tl-md rounded-tr-4xl rounded-br-md rounded-bl-4xl bg-white p-5">
                <motion.div
                    whileHover={{ y: -2 }}
                    className="h-11 w-11 rounded-md bg-zinc-200 p-3"
                >
                    <Barbell size={20} weight="fill" className="text-purple-500" />
                </motion.div>
                <div className="flex flex-col gap-0.5">
                    <h3
                        className="max-w-[8rem] truncate text-sm font-bold"
                        title={flow.title}
                    >
                        {flow.title}
                    </h3>
                    <p className="text-xs text-gray-500">{flow.level}</p>
                </div>
            </div>
        </motion.div>
    );
};

export default FlowItem;
