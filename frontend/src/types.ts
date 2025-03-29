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
