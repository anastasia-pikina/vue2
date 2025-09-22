import { createApp } from 'vue';
import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router';
import { createPinia } from 'pinia';
import App from './App.vue';
import Main from './Main.vue';
import News from './News.vue';
import NewsRouter from './NewsRouter.vue';
import Contacts from './Contacts.vue';
import NewDetail from './NewDetail.vue';
import NotFound from './NotFound.vue';
import './index.css';
import VueSmoothScroll from 'vue3-smooth-scroll';

const router = createRouter({
	routes: [
		{
			path: '/',
			component: Main,
			name: 'Main',
		},
		{
			path: '/news/',
			component: NewsRouter,
			children: [
				{
					path: 'page/:page?',
					component: News,
					name: 'NewsPage',
				},
				{
					path: '',
					component: News,
					name: 'News',
				},
				{
					path: ':id/',
					component: NewDetail,
					name: 'NewDetail',
				},
			],
		},
		{
			path: '/contacts/',
			component: Contacts,
			name: 'Contacts',
		},
		{
			path: '/:pathMatch(.*)*',
			component: NotFound
		},
	],
	history: createWebHistory(),
})

const app = createApp(App)
app.use(router)
app.use(VueSmoothScroll)
app.use(createPinia());
app.mount('#app')
