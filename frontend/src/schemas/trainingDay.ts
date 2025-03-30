import { z } from 'zod';

export const createNewTrainingDaySchema = z.object({
    title: z.string().min(1, 'Title is required'),
    day: z.enum([
        'Monday',
        'Tuesday',
        'Wednesday',
        'Thursday',
        'Friday',
        'Saturday',
        'Sunday',
    ]),
    duration: z.number().positive(),
});
