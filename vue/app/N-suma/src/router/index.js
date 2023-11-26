/**
 * router/index.js
 * included in `plugins/index.js`
 */

import { createRouter, createWebHistory } from "vue-router";
import { useStore } from "@/stores/user";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("@/layouts/Default.vue"),
      children: [
        {
          path: "",
          component: () => import("@/views/Index.vue"),
        },
      ],
    },
    {
      path: "/app",
      component: () => import("@/layouts/Default.vue"),
      children: [
        {
          path: "admin",
          component: () => import("@/layouts/Main.vue"),
          children: [
            {
              path: "",
              component: () => import("@/views/Admin.vue"),
            },
          ],
          // meta: { requiresAuth: true },
        },
      ],
      meta: { requiresAuth: true },
    },
    {
      path: "/:pathMatch(.*)*",
      redirect: "/",
    },
  ],
});

//appにアクセスしたときにログインしていなければログインページに飛ばす
router.beforeEach((to, from, next) => {
  const store = useStore();

  // storeのisLoginを取得。storeの情報を取得しきるまで待つ。
  const isLogin = store.isLogin;

  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  if (requiresAuth && !isLogin) {
    console.log(isLogin);
    alert("ログインしてください");
    next({ path: "/" });
    return;
  } else {
    next();
    return;
  }
});

export default router;
