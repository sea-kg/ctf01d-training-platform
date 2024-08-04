import { createApp } from 'vue';
import "./style.scss";
import App from './App.vue';
import router from './router/router';
// import axios from 'axios';
import axios from 'https://cdn.jsdelivr.net/npm/axios@1.3.5/+esm';

const app = createApp(App);

app.config.globalProperties.$axios = axios;

app.use(router).mount('#app');
