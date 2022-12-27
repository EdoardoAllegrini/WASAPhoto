import {createRouter, createWebHashHistory} from 'vue-router'
import LoginPage from '../views/LoginPage.vue'
import Stream from '../views/Stream.vue'
import Profile from '../components/Profile.vue'
import PageNotFound from '../components/PageNotFound.vue'
import Image from '../components/image.vue'
import FlwPopup from '../components/flwPopup.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: LoginPage},
		{path: '/', component: LoginPage},
		// {path: '/stream/', component: Stream},
		{path: '/stream', component: Stream, children: [{path: '/users/:username/media/:photo', component: Image}]},
		{path: '/users/:username', component: Profile, children: [{path: '/users/:username/followers', component: FlwPopup, props: {recv: Object}}, {path: '/users/:username/following', component: FlwPopup, props: {recv: Object}}]},
		// {path: '/users/:username/media/:photo', component: Image},
		// {path: '/users/:username/followers', component: FlwPopup},
		// {path: '/users/:username/following', component: FlwPopup},
		{path: '/:pathMatch(.*)*', component: PageNotFound},
	]
})

export default router
