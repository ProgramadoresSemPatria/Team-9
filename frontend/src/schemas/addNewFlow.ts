import { z } from 'zod';

export const addNewFlowSchema = z.object({
    title: z
        .string()
        .trim()
        .min(1, 'The title must have at least 1 character')
        .max(50, 'The title can not exceed 50 characters'),
    workoutLevel: z.enum(['beginner', 'intermediate', 'advanced']),
    cover: z.string().optional(),
});
