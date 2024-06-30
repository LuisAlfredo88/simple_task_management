import type { TaskContract } from '../../domain/contract/taskContract'
import type { Task, TaskFilter, TaskRecord, TaskStatus } from '../../domain/entity/task'
import { post, get, patch, DEFAULT_API_PATH, deleteRecord } from '@/modules/shared/http_handler'
import { getTaskRecords, getTaskSingleRecords } from '../adapters/taskRecords'

type RequestResponse = {
    status: number;
    data: any;
    headers: any;
}

const loadTasks = async (filter: TaskFilter | null): Promise<GridRecord> => {
	const response = {
		status: 200,
		data: [
			{
				id: 1,
				title: "Task #1 from admin",
				description: "This is the first registered task",
				status: "to_do",
				color: "#03A9F4",
				createdBy: "Admin Main",
				createdById: "7414a010-b098-448c-8dd0-898c495bd9d6",
				assignedTo: "Admin Main"
			},	
			{	
				id: 2,
				title: "Task #2 from admin",
				description: "This is the second registered task",
				status: "in_progress",
				color: "#FF9800",
				createdBy: "Admin Main",
				createdById: "7414a010-b098-448c-8dd0-898c495bd9d6",
				assignedTo: "Admin Main"
			},
		],
		headers: {
			'x-totalrecords': 50
		}
	} as RequestResponse;

	return getTaskRecords(response);
}

const saveTask = async (task: Task): Promise<TaskRecord | null> => {
	const response = {
		status: 200,
		data: {
			id: 1,
			title: "Task #1 from main user",
			description: "This is the first registered task",
			status: "to_do",
			color: "#03A9F4",
			createdBy: "Main First",
			createdById: "7414a010-b098-448c-8dd0-898c495bd9d6",
			assignedTo: ""
		}
	} as RequestResponse;

	return getTaskSingleRecords(response);
}

const getTaskById = async (taskId: number): Promise<Task> => {
	return {
		id: taskId,
		title: "task title"
	} as Task
}

const changeTaskStatus = async (taskId: number, status: number): Promise<boolean> => {
	const requestData = await patch(DEFAULT_API_PATH + `/tasks_management/tasks/${taskId}/${status ? 1 : 0}`, {} );
	return requestData.status === 200;
}

const taskExists = async (taskName: string): Promise<boolean> => {
	const requestData = await get(DEFAULT_API_PATH + `/tasks_management/tasks/${taskName}/exists`, {} );
	return requestData.data.exists;
}

const getTaskStatus = async (): Promise<TaskStatus[]> => {
	const requestData = await get(DEFAULT_API_PATH + `/tasks_management/task_status`, {} );
	return requestData.data;
}

const deleteTask = async (taskId: number): Promise<boolean> => {
	return true;
}

export default {
	loadTasks,
	saveTask,
	getTaskById,
	changeTaskStatus,
	taskExists,
	getTaskStatus,
	deleteTask
} as TaskContract