import TaskRepository from '../repository/taskRepository'

export default [
    {
        path:'/tasks/:id?/:dialogType?',
        name:'tasks',
        component: () => import('../ui/task/Task.vue'),
        props: {
            TaskRepository,
        }
    },
];
