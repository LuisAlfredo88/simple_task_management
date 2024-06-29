import type { User, UserForm } from "../entity/user"

export interface UserContract {
    saveUser(user: UserForm): Promise<User>
    getAllUsers(filter: CriteriaFilter | null): Promise<GridRecord>,
    getUserById(userId: string):Promise<UserForm>,   
    changeUserStatus(userId: string, status: boolean): Promise<boolean>   
    loadUsers(): Promise<User[]>,
    userExists(userName: string): Promise<boolean>   
}
