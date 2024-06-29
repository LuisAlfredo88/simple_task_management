import axios from "axios";
import { useAuthStore } from "@/modules/security/store/auth";
import { useAuth } from "@/modules/security/composable/useAuth";

export const DEFAULT_API_PATH = import.meta.env.VITE_API_URL;

export const get = (url: string, params: Record<string, unknown> = {}): Promise<any> => {
    return axios.get<Response>(url, { params });
}

export const post = (url: string, params: Record<string, unknown>): Promise<any> => {
    return axios.post(url, params);
}

export const put = (url: string, data: Record<string, unknown>): Promise<any> => {
    return axios.put(url, data);
}

export const patch = (url: string, data: Record<string, unknown>): Promise<any> => {
    return axios.patch(url, data);
}

export const deleteRecord = (url: string): Promise<any> => {
    return axios.delete(url);
}

axios.interceptors.request.use((config) => {
    const system = useAuthStore();
    const user_token = system.token;
	if(!user_token) return config;
    
    config.headers.Authorization =  `Bearer ${user_token}`;
    return config;
}, (error) => {
	return Promise.reject(error);
});

axios.interceptors.response.use(
    response => response,
    error => {
        // Forcing user to logout if any request response code is 401
        if(error.response.status === 401) {
            const auth = useAuth();
            auth.logout();
        }
    });