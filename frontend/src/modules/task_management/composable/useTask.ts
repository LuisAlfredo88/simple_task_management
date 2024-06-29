import { getErrorInformation } from '@/modules/shared/domain/errors'
import { systemMessage } from '@/modules/system/domain/events/systemMessage'
import { 
    LOADING_DATA_ERROR, 
    SAVING_RECORD_ERROR,
    DELETE_RECORD_ERROR,
    DELETE_RECORD_SUCCESS
} from '@/modules/shared/domain/commonMessages'
import type { TaskContract } from '../domain/contract/taskContract'
import type { Task, TaskFilter, TaskStatus, TaskRecord } from '../domain/entity/task'
import { validateTask} from '../domain/entity/task'
import { evaluateFilter, resetPageScroll } from '@/modules/shared/utility/records'
import { ref } from 'vue'

const tasks = ref<TaskRecord[]>([]);
const totalRecords = ref(0);

const filter = ref({
    filters: {} as TaskFilter,
    limit: 50,
    skip: 0
} as CriteriaFilter);

export type taskProps = {
    taskRepository : TaskContract,
    taskId?: string
}

export const useTask = (props: taskProps) => {
    const lastFilterHash = ref(0);
    const loadingRecords = ref(false);
    const taskId = ref('');
    const savingRecord = ref(false);
    const showFilter = ref(false);
    const showDialog = ref(false);
    const taskForm = ref({} as Task);
    const taskAlreadyExists = ref(false);

    const setTaskData = (task: Task): void => {
        taskForm.value = task;
    }

    const getAllTasks = async (): Promise<TaskRecord[]> => {    
        const hasChanges = evaluateFilter(filter, lastFilterHash);
        
        if(totalRecords.value > 0 && totalRecords.value <= (filter.value.skip || 0)) {
            return tasks.value;
        }
        
        loadingRecords.value = true;

        try {
            const data = await props.taskRepository.loadTasks(filter.value);
            if(hasChanges) {
                resetPageScroll();
                tasks.value = [...(data.records || [])];
            } else {
                tasks.value = [...tasks.value, ...(data.records || [])];
            }
            totalRecords.value = data.totalRecordsQty;
            
        } catch(e) {
            systemMessage({ "type": "error", "description": LOADING_DATA_ERROR });
        }
        
        loadingRecords.value = false;

        return tasks.value;
    }

    const saveTask = async (): Promise<Task | null> => {
        try {
            validateTask(taskForm.value);

            const task = await props.taskRepository.saveTask(taskForm.value);
            if(!task) return null;

            if(taskForm.value.id) { // Updating record if is an update
                const index = tasks.value.findIndex(x => x.id === taskForm.value.id);
                tasks.value[index] = task;
            } else {                
                tasks.value = [task, ...tasks.value];
                totalRecords.value += 1;
            }
            
            resetPageScroll();
            
            return task;
        } catch(e) {
            const error = getErrorInformation(e as Error, SAVING_RECORD_ERROR);
            systemMessage({ "type": error.type, "description": error.message });
        }

        return null;
    }

    const getTaskById = async (taskId: number): Promise<Task | null> => {
        try {
            return await props.taskRepository.getTaskById(taskId);            
        } catch(e) {
            systemMessage({ "type": "error", "description": LOADING_DATA_ERROR });
        }

        return null;
    }

    const changeTaskStatus = async (taskId: number, taskStatus: number): Promise<boolean> => {        
        const result = await props.taskRepository.changeTaskStatus(taskId, taskStatus);
        return result
    }
    
    const getTaskStatus = async (): Promise<TaskStatus[]> => {  
        try {
            return await props.taskRepository.getTaskStatus();          
        } catch(e) {
            systemMessage({ "type": "error", "description": LOADING_DATA_ERROR });
        }

        return [];
    }

    const deleteTask = async (taskId: number): Promise<void> => {

        try {
            const sucess =  await props.taskRepository.deleteTask(taskId);   
            if(!sucess) {
                systemMessage({ "type": "error", "description": DELETE_RECORD_ERROR });
                return;
            }
        } catch(x) {
            systemMessage({ "type": "error", "description": DELETE_RECORD_ERROR });
            return;
        }

        tasks.value = tasks.value.filter(x => x.id !== taskId);

        systemMessage({ "type": "success", "description": DELETE_RECORD_SUCCESS });
        
        return;
    }

    return {
        totalRecords,
        tasks,
        taskForm,
        loadingRecords,
        filter,
        taskId,
        savingRecord,
        showFilter,
        showDialog,
        taskAlreadyExists,
        setTaskData,
        getAllTasks,
        saveTask,
        getTaskById,
        changeTaskStatus,
        getTaskStatus,
        deleteTask
    }
}

