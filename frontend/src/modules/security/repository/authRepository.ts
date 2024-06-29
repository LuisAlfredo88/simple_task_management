import { post, DEFAULT_API_PATH } from "@/modules/shared/http_handler"
import type { AuthContract } from "../domain/contract/AuthContract"
import Menu from './static_data/menu.json'

const authenticate = async (userName: string, password: string): Promise<any> => {
    const data = await post(DEFAULT_API_PATH + '/security/authenticate', {
        userName, password
    });

    const loginInfo = {
        ...data.data,
        menu: Menu,
    }
    return loginInfo;
}

const refreshToken = async (): Promise<string> => {
    const response = await post(DEFAULT_API_PATH + '/security/token_refresh', { });
    return response.data.token;
}

export default {
    authenticate,
    refreshToken
} as AuthContract
