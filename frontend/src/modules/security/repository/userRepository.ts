import type { UserRepository as UserRepo } from '../domain/contract/userContract'
import type { User, UserForm, UserFilter } from '../domain/entity/user'
import { post, get, patch, DEFAULT_API_PATH } from '@/modules/shared/http_handler'
import { getUserRecords, getUserSingleRecords } from './adapters/userRecords'

const loadUsers = async (filter: UserFilter | null): Promise<GridRecord> => {
    const queryFilter = {
        limit: filter?.limit,
        skip: filter?.skip,
        search: filter?.filters.search,
        status: (filter?.filters?.status === undefined || filter?.filters?.status === null)  
                ? null : (filter?.filters?.status ? 1: 0),
    }

    const response = await get(DEFAULT_API_PATH + '/security/users', queryFilter );
    return getUserRecords(response);
}

const saveUser = async (user: UserForm): Promise<User> => {
    const requestData = await post(DEFAULT_API_PATH + '/security/users', user);
    return getUserSingleRecords(requestData);
}

const getUserById = async (userId: string): Promise<User> => {
    const requestData = await get(DEFAULT_API_PATH + '/security/users/' + userId);
    return requestData.data;
}

const changeUserStatus = async (userId: string, status: boolean): Promise<boolean> => {
    const requestData = await patch(DEFAULT_API_PATH + `/security/users/${userId}/${status ? 1 : 0}`, {} );
    return requestData.status === 200;
}

const userExists = async (userName: string): Promise<boolean> => {
    const requestData = await get(DEFAULT_API_PATH + `/security/users/${userName}/exists`, {} );
    return requestData.data.exists;
}

export default {
    loadUsers,
    saveUser,
    getUserById,
    changeUserStatus,
    userExists
} as UserRepo