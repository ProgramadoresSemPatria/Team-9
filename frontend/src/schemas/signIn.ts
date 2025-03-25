import { z } from 'zod';

export const signInSchema = z.object({
    email: z.string().trim().email('Invalid email format'),
    password: z.string().trim().min(1, 'The password is obrigatory'),
});
