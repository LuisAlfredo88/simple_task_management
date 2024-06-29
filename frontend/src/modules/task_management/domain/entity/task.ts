import { validator } from "@/modules/shared/domain/errors";

export type Task  = {
    id: number;
    title: string;
    description: string;
    statusId: number;
    assignedTo: string;   
}

export type TaskRecord  = {
    id: number;
    title: string;
    description: string;
    statusId: number;   
    status: string;
    color: string;
    createdBy: string; 
    assignedTo: string;     
}

export type TaskStatus = {
    id: number;
    name: string;
    color: string;
}

export type TaskFilter = CriteriaFilter & {
    search?: string;
    statusId: number; 
    createdBy: string; 
    assignedTo: string; 
}

export const validateTask = (task: Task) => {
    const dataValidation = [
        { validation: !task.title, error: 'TASK.must_specify_title', tag: 'title' },
        { validation: !task.description, error: 'TASK.must_specify_description', tag: 'description' },
        { validation: !task.statusId, error: 'TASK.must_specify_status', tag: 'status' },
    ];

    validator(dataValidation);
}
