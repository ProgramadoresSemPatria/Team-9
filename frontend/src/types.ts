export type Exercise = {
    id: string;
    title: string;
    muscle_group: string;
    repetitions: number;
    sets: number;
};

export type TrainingDay = {
    id?: string;
    title: string;
    dayOfWeek: string;
    exercises: number; //Exercise[]
    duration: number;
};

export type Flow = {
    id: string;
    title: string;
    level: string;
};

export type RegisterUser = {
    name: string;
    email: string;
    password: string;
};

export type Login = {
    email: string;
    password: string;
};

export type CreateExercise = {
    title: string;
    muscle_group: string;
    repetitions: number;
    sets: number;
};

export type CreateTrainingDay = {
    title: string;
    day: string;
    duration: number;
};

export type CreateFlow = {
    title: string;
    level: 'beginner' | 'intermediate' | 'advanced';
    cover?: File;
};
