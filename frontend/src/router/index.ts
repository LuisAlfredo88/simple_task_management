import { createRouter, createWebHashHistory } from 'vue-router'
import MainView from '../layouts/main/Index.vue'
import Login from '@/modules/security/ui/login/Login.vue'

import authRepository from '@/modules/security/repository/authRepository'
import { useSecurity } from '@/modules/security/composable/useSecurity'
import SecurityRouter from '@/modules/security/router/securityRouter'

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: '/login',
			name: 'login',
			component: Login,
			props: {
				authRepository
			}
		},

		{
			path: '/',
			name: 'main',
			component: MainView,
			children: [
				{
					path: '/',
					name: 'home',
					component: () => import('../views/Home.vue')
				},
			]
		},

		{
			path: '/security',
			name: 'security',
			component: MainView,
			children: [
				...SecurityRouter		
			]
		},
	]
});

const setRouterSecurity = () => {
	router.beforeEach((to, from, next) => {
		const { userIsLogged } = useSecurity();

		if(!userIsLogged.value && (to.path != '/login'))
			next('/login');        
		else{
			if(userIsLogged.value && to.path == '/login') {
				next('/');
				return;
			}
	
			next();
		}
	
		window.scrollTo(0, 0);
	});
};

export const useRouter = ((app: any) => {
	app.use(router);
	setRouterSecurity();
});

export default router;
