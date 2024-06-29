import './assets/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import { useRouter } from './router'
import usePrimeVue from './plugins/prime-vue'
import useInternalization from './plugins/i18n'
import useDatePipe from './plugins/pipes/date'

const app = createApp(App);
app.use(createPinia());

usePrimeVue(app);
useInternalization(app);
useRouter(app);
useDatePipe(app);
app.mount('#app');
