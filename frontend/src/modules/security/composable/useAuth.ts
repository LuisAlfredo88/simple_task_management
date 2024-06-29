import type { AuthContract } from '../domain/contract/authContract'
import { getEncrypt, setEncrypt, unset as unsetStore } from '@/modules/shared/utility/localStorage'
import type { Auth } from '../domain/entity/auth'
import { ref } from 'vue'
import { useAuthStore } from '../store/auth'
import { useRouter } from 'vue-router'

const  SESSION_NAME = 'session';

const setUserData = (sytemStore: any, authData: Auth): void => {
    sytemStore.setMenu(authData.menu);
    sytemStore.setToken(authData.token);
    sytemStore.setUser(authData.user);
    setEncrypt(SESSION_NAME, authData, true);
}

const clearUserData = (sytemStore: any): void => {
    sytemStore.setMenu([]);
    sytemStore.setToken(null);
    sytemStore.setUser(null);
    unsetStore(SESSION_NAME);
}

export const useAuth = (authRepo?: AuthContract) => {
    const loginError = ref(false);
    const user = ref('');
    const password = ref('');
    const sytemStore = useAuthStore();
    const router = useRouter();

    const login = async() => {

        if(!authRepo) {
            throw new Error('must specify auth repository');
        }
        
        try {
            const authData = await authRepo.authenticate(user.value, password.value);
            setUserData(sytemStore, authData);
            router.push('/');
            scheduleTokenRenew();
        } catch(error) {
            loginError.value = true;
        }
    }

    const logout = () => {
        clearUserData(sytemStore);
        window.location.href = '#/';
    }

    const scheduleTokenRenew = () => {
        const interval = 5; // minutes

        const refreshToken = async () => {
            const token = await authRepo?.refreshToken();
            const session = getEncrypt('session', true) as Auth;
            session.token = token || '';
            setUserData(sytemStore, session);
        }

        refreshToken();
        setInterval(() => {
            refreshToken();
        }, interval * 60 * 1000);
    }

    return {
        loginError,
        user,
        password,
        login,
        logout,
        scheduleTokenRenew
    }
}
