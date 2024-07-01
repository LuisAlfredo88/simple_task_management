import type { TaskContract } from '../../domain/contract/taskContract'
import type { Task, TaskFilter, TaskRecord } from '../../domain/entity/task'
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
	return true;
}

const deleteTask = async (taskId: number): Promise<boolean> => {
	return true;
}

export default {
	loadTasks,
	saveTask,
	getTaskById,
	changeTaskStatus,
	deleteTask
} as TaskContract