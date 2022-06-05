import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { onAuthStateChanged, signOut } from "firebase/auth"
import TopView from '../views/TopView.vue'
import { FirebaseAuth } from '@/main'

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
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('../views/HomeView.vue'),
    meta: {
      requireAuth: true,
    }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((rec) => rec.meta.requireAuth)) {
    // 認証チェック
    await onAuthStateChanged(FirebaseAuth, async (user) => {
      if (user !== null) {
        if (to.name === "SignIn") {
          //
        }
        next()
      } else {
        await signOut(FirebaseAuth)
        router.replace({ name: "SignIn" })
      }
    })
  } else {
    next()
  }
})

export default router
