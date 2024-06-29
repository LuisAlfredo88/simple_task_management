import type { AuthRepository } from "../domain/repository/authRepository"
import type { Auth } from "../domain/entity/auth"
import type { User } from "../domain/entity/user"
import Menu from './static_data/menu.json'

export class MockAuthRepository implements AuthRepository {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    authenticate(userName: string, password: string): Promise<Auth> {
        return new Promise((resolve) => {
            const user: User = {
                userId: 1,
                userName: 'Admin'
            };

            const authData: Auth =  {
                user,
                privileges: ['create_user'],
                token: 'yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtaXNhZWRpZm8iLCJqdGkiOiJlYzZjZmZjYS0wZTU5LTQ4YzMtYjlhYy0zNGMyOTA4YTY4YzEiLCJlbWFpbCI6Im1pc2FlZGlmb0BnbWFpbC5jb20iLCJyb2xlIjoiQmFucXVlcm8iLCJ1aWQiOiI0YTNlM2Z',
                menu: Menu
            };

            resolve(authData);
        })
    }
}
