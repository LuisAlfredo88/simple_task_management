import type { User } from "../../domain/entity/user";

export const getUserRecords = (response: RequestResponse): GridRecord => {
    return {
        records: response.data,
        totalRecordsQty: Number(response.headers['x-totalrecords'])
    }
}

export const getUserSingleRecords = (response: RequestResponse): User => {
    const user = response.data;
    return user;
}