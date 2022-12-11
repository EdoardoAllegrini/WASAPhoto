import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginPage from '../views/LoginPage.vue'
import Stream from '../views/Stream.vue'
import Profile from '../components/Profile.vue'
import PageNotFound from '../components/PageNotFound.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: LoginPage},
		{path: '/', component: LoginPage},
		{path: '/stream/:username', component: Stream},
		{path: '/users/:username', component: Profile},
		{path: '/:pathMatch(.*)*', component: PageNotFound},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
