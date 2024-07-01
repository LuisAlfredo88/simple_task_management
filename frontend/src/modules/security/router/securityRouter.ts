import UserRepository from '../repository/userRepository'

export default [
    {
        path:'/users/:id?/:dialogType?',
        name:'users',
        component: () => import('../ui/user/User.vue'),
        props: {
            UserRepository,
        }
    },
];
