import { z } from 'zod';

export const registerSchema = z
    .object({
        name: z
            .string()
            .trim()
            .min(2, 'The name must be at least 2 characters long')
            .max(50, 'The name cannot exceed 50 characters'),

        email: z.string().trim().email('Invalid email format'),

        password: z
            .string()
            .trim()
            .min(8, 'The password must be at least 8 characters long')
            .regex(
                /[A-Z]/,
                'The password must contain at least one uppercase letter'
            )
            .regex(
                /[a-z]/,
                'The password must contain at least one lowercase letter'
            )
            .regex(/[0-9]/, 'The password must contain at least one number')
            .regex(
                /[@$!%*?&]/,
                'The password must contain at least one special character (@, $, !, %, *, ?, &)'
            )
            .max(100, 'The password cannot exceed 100 characters'),

        confirmPassword: z.string(),
    })
    .superRefine((data, ctx) => {
        if (data.confirmPassword !== data.password) {
            ctx.addIssue({
                code: 'custom',
                path: ['confirmPassword'],
                message: 'Passwords do not match',
            });
        }
    });
