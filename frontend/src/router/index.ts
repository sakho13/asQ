import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { onAuthStateChanged, signOut } from "firebase/auth";

import { FirebaseAuth } from "@/main";
import TopView from "../views/TopView.vue";
import { userStore } from "../pinia/userStore";
import { Api } from "@/apis/Api";

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
  {
    path: "/about",
    name: "About",
    component: () => import("../views/AboutView.vue"),
  },
  {
    path: "/signout",
    name: "SignOut",
    component: () => import("../views/SignOutView.vue"),
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
    onAuthStateChanged(FirebaseAuth, async (user) => {
      if (user !== null) {
        const res = await Api.createUser({
          firebase_jwt: user.uid,
        });
        if (res.result_flg === 1) {
          return next();
        } else {
          return next({ name: "Top" });
        }
      } else {
        await signOut(FirebaseAuth);
        return next({ name: "Top" });
      }
    });
  } else {
    return next();
  }
});

export default router;
