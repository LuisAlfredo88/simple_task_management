import { ref } from 'vue'
import { defineStore } from 'pinia'
import { getEncrypt } from '@/modules/shared/utility/localStorage'
import type { Auth } from '../domain/entity/auth';
import type { User } from '../domain/entity/user';

export const useAuthStore = defineStore('auth', () => {
    const session = getEncrypt('session', true) as Auth;

    const token = ref(session?.token);
    const systemMenu = ref( (session?.menu || []) as Record<string, any>[])
    const user = ref(session?.user || {});
  
    function setMenu(menu: Record<string, any>[]) {
        systemMenu.value = menu;
    }

    function setToken(authToken: string) {
        token.value = authToken;
    }

    function setUser(logguedUser: User) {
        user.value = logguedUser;
    }

    return { 
        systemMenu, token, user,
        setMenu, setToken, setUser
    }
})
