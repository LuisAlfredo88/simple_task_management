import type { TaskContract } from '../domain/contract/taskContract'
import type { Task, TaskFilter, TaskRecord, TaskStatus } from '../domain/entity/task'
import { post, get, patch, DEFAULT_API_PATH } from '@/modules/shared/http_handler'
import { getTaskRecords, getTaskSingleRecords } from './adapters/taskRecords'

const loadTasks = async (filter: TaskFilter | null): Promise<GridRecord> => {
	const queryFilter = {
		limit: filter?.limit,
		skip: filter?.skip,
		search: filter?.filters.search,
		createdBy: (filter?.filters?.createdBy === undefined || filter?.filters?.createdBy === null)  
				? null : filter.filters.createdBy,
		assignedTo: (filter?.filters?.assignedTo === undefined || filter?.filters?.assignedTo === null)  
				? null : filter.filters.assignedTo,                             
		status: (filter?.filters?.statusId === undefined || filter?.filters?.statusId === null)  
				? null : filter.filters.statusId
	}

	const response = await get(DEFAULT_API_PATH + '/tasks_management/tasks', queryFilter );
	return getTaskRecords(response);
}

const saveTask = async (task: Task): Promise<TaskRecord | null> => {
	const requestData = await post(DEFAULT_API_PATH + '/tasks_management/tasks', task);
	return getTaskSingleRecords(requestData);
}

const getTaskById = async (taskId: number): Promise<Task> => {
	const requestData = await get(DEFAULT_API_PATH + '/tasks_management/tasks/' + taskId);
	return requestData.data;
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

export default {
	loadTasks,
	saveTask,
	getTaskById,
	changeTaskStatus,
	taskExists,
	getTaskStatus
} as TaskContract