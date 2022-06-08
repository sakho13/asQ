import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { onAuthStateChanged, signOut } from "firebase/auth";

import { FirebaseAuth } from "@/main";
import TopView from "../views/TopView.vue";
import { userStore } from "../pinia/userStore";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Top",
    component: TopView,
  },
  {
    path: "/sign_in",
    name: "SignIn",
    component: () => import("../views/SigninView.vue"),
  },
  {
    path: "/home",
    name: "Home",
    component: () => import("../views/HomeView.vue"),
    meta: {
      requireAuth: true,
    },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const userStoreObj = userStore();

  if (to.name === "Home") {
    if (userStoreObj.jwt !== "") {
      return next();
    } else {
      return next({ name: "SignIn" });
    }
  }

  if (to.matched.some((rec) => rec.meta.requireAuth)) {
    // 認証チェック
    await onAuthStateChanged(FirebaseAuth, async (user) => {
      if (user !== null) {
        return next();
      } else {
        await signOut(FirebaseAuth);
        return next({ name: "SignIn" });
      }
    });
  } else {
    return next();
  }
});

export default router;
