import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App2 from './App2.vue'
import App from './App.vue'
import News from './News.vue'
import NewsRouter from './NewsRouter.vue'
import Contacts from './Contacts.vue'
import NewDetail from './NewDetail.vue'
import './index.css'
import VueSmoothScroll from 'vue3-smooth-scroll'

const router = createRouter({
	routes: [
		{
			path: '/',
			component: App
		},
		{
			path: '/news/',
			component: NewsRouter,
			children: [
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

const app = createApp(App2)
app.use(router)
app.use(VueSmoothScroll)
app.mount('#app')
