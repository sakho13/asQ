import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import TopView from '../views/TopView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Top',
    component: TopView,
  },
  {
    path: '/sign_up',
    name: 'SignUp',
    component: () => import('../views/SignupView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
