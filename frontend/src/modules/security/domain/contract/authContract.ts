import type { Auth } from "../entity/auth"

export interface AuthContract {
    authenticate(userName: string, password: string): Promise<Auth>
    refreshToken(): Promise<string>
}
