import { expect, test } from 'vitest'
import { useTask } from './useTask'
import mockTaskRepository from '../repository/mock/mockTaskRepository'
import type { Task, TaskRecord } from '../domain/entity/task'
import { registerEvent } from '@/modules/shared/utility/eventHandler'

const props = {
    taskRepository: mockTaskRepository
};

let composable: ReturnType<typeof useTask>;

const setup = () => {
    composable = useTask(props);
    registerEvent('SYSTEM_MESSAGE', ()=>{}, true);
}

const mockTask = {
    id: 1, 
    title: 'task title' 
} as Task;

const mockTaskRecord = {
    id: 1,
    title: 'task title',
    assignedTo: '28928392892-478738437-34533',
    statusId: 1,
    description: 'Task Description'
} as Task;

const taskRecord = {
    id: 1,
    title: 'task title',
    description: 'Task Description',
    assignedTo: '28928392892-478738437-34533',
} as TaskRecord;

test('getAllTasks should fetch tasks successfully', async () => {
    composable = useTask(props);
    await composable.getAllTasks();
    
    expect(composable.tasks.value.length).toEqual(2);
    expect(composable.totalRecords.value).toBe(50);
    expect(composable.loadingRecords.value).toBe(false);
});

test('TaskData for saveTask method should be valid', async () => {
    composable = useTask(props);
    const record = JSON.parse(JSON.stringify(mockTaskRecord)); 
    record.title = '';
    record.description = '';
    record.statusId = 0;

    composable.taskForm.value = record;
    await expect(() => composable.saveTask()).rejects.toThrowError()
});

test('saveTask should be run successfully', async () => {
    setup();
    const record = JSON.parse(JSON.stringify(mockTaskRecord));
    record.id = 0;
    composable.taskForm.value = record;
    let taskCount = composable.tasks.value.length;
    await composable.saveTask();
    const taskCountAfeterInsert = composable.tasks.value.length;

    // When it is a new record
    expect(taskCountAfeterInsert).toEqual(taskCount + 1);

    record.id = 1;
    composable.taskForm.value = record;
    taskCount = composable.tasks.value.length;
    await composable.saveTask();
    const taskCountAfeterUpdate = composable.tasks.value.length;
    // When it is an existing record
    expect(taskCountAfeterUpdate).toEqual(taskCount);
});

test('getTaskById should fetch a task by ID', async () => {
    setup();
    const taskId = 1;
    mockTaskRepository.getTaskById(taskId);
    const result = await composable.getTaskById(taskId);
    expect(result).toEqual(mockTask);
});

test('deleteTask should delete a task successfully', async () => {
    setup();
    composable = useTask(props);
    const taskId = 1;
    composable.tasks.value.push(taskRecord);
    await composable.deleteTask(taskId);
    expect(composable.tasks.value.some(task => task.id === taskId)).toBe(false);
});
