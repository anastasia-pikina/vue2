import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import { createPinia } from 'pinia';
import App from './App.vue';
import Main from './Main.vue';
import News from './News.vue';
import NewsRouter from './NewsRouter.vue';
import Contacts from './Contacts.vue';
import NewDetail from './NewDetail.vue';
import './index.css';
import VueSmoothScroll from 'vue3-smooth-scroll';

const router = createRouter({
	routes: [
		{
			path: '/',
			component: Main
		},
		{
			path: '/news/',
			component: NewsRouter,
			children: [
				{
					path: 'page/:page?',
					component: News,
					name: 'news',
				},
				{
					path: '',
					component: News,
				},
				{
					path: ':id/',
					component: NewDetail,
				},
			],
		},
		{
			path: '/contacts/',
			component: Contacts
		},
	],
	history: createWebHistory()
})

const app = createApp(App)
app.use(router)
app.use(VueSmoothScroll)
app.use(createPinia());
app.mount('#app')
