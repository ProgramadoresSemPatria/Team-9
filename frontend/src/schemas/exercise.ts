import { z } from 'zod';

export const createExerciseSchema = z.object({
    title: z.string().min(1, 'Title is required'),
    muscle_group: z.string().min(1, 'Muscle is required'),
    repetitions: z.number().int().positive('Repetitions must be a positive integer'),
    sets: z.number().int().positive('Sets must be a positive integer'),
});
