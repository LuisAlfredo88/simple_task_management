import UserRepository from '../repository/userRepository'

export default [
    {
        path:'/users/:id?/:dialogType?',
        name:'users',
        component: () => import('../ui/user/User.vue'),
        meta: {
            i18n: ['security']
        },
        props: {
            UserRepository,
        }
    },
];
