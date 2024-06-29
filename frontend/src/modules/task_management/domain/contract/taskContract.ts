import type { Task, TaskStatus, TaskRecord } from "../entity/task"

export interface TaskContract {
    saveTask(task: Task): Promise<TaskRecord | null>
    loadTasks(filter: CriteriaFilter | null): Promise<GridRecord>,
    getTaskById(taskId: number):Promise<Task>,       
    changeTaskStatus(taskId: number, status: number): Promise<boolean>    
    getTaskStatus(): Promise<TaskStatus[]>    
    deleteTask(taskId: number): Promise<boolean>    
}
