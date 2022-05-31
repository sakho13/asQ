import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import TopView from '../views/TopView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Top',
    component: TopView,
  },
  {
    path: '/sign_in',
    name: 'SignIn',
    component: () => import('../views/SignInView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
