import type { AuthContract } from "../domain/contract/authContract"
import type { Auth } from "../domain/entity/auth"
import type { User } from "../domain/entity/user"
import Menu from './static_data/menu.json'

const authenticate = async (userName: string, password: string): Promise<any> => {

    const user: User = {
        id: 'yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9',
        userName: 'admin',
        name: "Admin",
        lastName: "Main",
        isActive: true,
        password: "123456",
    };

    const authData: Auth =  {
        user,
        token: 'yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtaXNhZWRpZm8iLCJqdGkiOiJlYzZjZmZjYS0wZTU5LTQ4YzMtYjlhYy0zNGMyOTA4YTY4YzEiLCJlbWFpbCI6Im1pc2FlZGlmb0BnbWFpbC5jb20iLCJyb2xlIjoiQmFucXVlcm8iLCJ1aWQiOiI0YTNlM2Z',
        menu: Menu
    };

    return authData;
}

const refreshToken = async (): Promise<string> => {
    return 'yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtaXNhZWRpZm8iLCJqdGkiOiJlYzZjZmZjYS0wZTU5LTQ4YzMtYjlhYy0zNGMyOTA4YTY4YzEiLCJlbWFpbCI6Im1pc2FlZGlmb0BnbWFpbC5jb20iLCJyb2xlIjoiQmFucXVlcm8iLCJ1aWQiOiI0YTNlM2Z';
}

export default {
    authenticate,
    refreshToken
} as AuthContract

