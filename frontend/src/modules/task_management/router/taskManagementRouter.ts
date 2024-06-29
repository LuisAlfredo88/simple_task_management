import TaskRepository from '../repository/taskRepository'
import UserRepository from '../../security/repository/userRepository'

export default [
    {
        path:'/tasks/:id?/:dialogType?',
        name:'tasks',
        component: () => import('../ui/task/Task.vue'),
        props: {
            TaskRepository,
            UserRepository
        }
    },
];
