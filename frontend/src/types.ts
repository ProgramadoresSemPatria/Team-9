export type Exercise = {
    id?: string;
    title: string;
    muscle: string;
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
    id?: string;
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
    trainingDayId: string;
    title: string;
    muscle: string;
    repetitions: number;
    sets: number;
};

export type CreateTrainingDay = {
    flowId: string;
    title: string;
    dayOfWeek: string;
    exercises: Exercise[];
    duration: number;
};

export type CreateFlow = {
    title: string;
    level: 'beginner' | 'intermediate' | 'advanced';
    cover?: File;
};
