import { useAuthStore } from '../store/auth'
import { computed } from 'vue'

export const useSecurity = () => {
    const system = useAuthStore();
    const userIsLogged = computed(() => (system?.token !== undefined && system?.token !== null));
    const user = computed(() => system.user);

    return {
        userIsLogged,
		user
    }
}
