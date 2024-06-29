import type { TaskRecord } from "../../domain/entity/task";

export const getTaskRecords = (response: RequestResponse): GridRecord => {
    return {
        records: response.data,
        totalRecordsQty: Number(response.headers['x-totalrecords'])
    }
}

export const getTaskSingleRecords = (response: RequestResponse): TaskRecord => {
    const task = response.data;
    return task;
}