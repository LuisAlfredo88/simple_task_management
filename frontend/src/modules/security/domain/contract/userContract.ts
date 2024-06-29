import type { User, UserForm } from "../entity/user"

export interface UserRepository {
    saveUser(user: UserForm): Promise<User>
    loadUsers(filter: CriteriaFilter | null): Promise<GridRecord>,
    getUserById(userId: string):Promise<UserForm>,   
    changeUserStatus(userId: string, status: boolean): Promise<boolean>   
    userExists(userName: string): Promise<boolean>   
}
