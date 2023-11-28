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
router.beforeEach(async (to, from, next) => {
  const store = useStore();

  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
  const requiresAdmin = to.matched.some((record) => record.meta.requiresAdmin);
  const requiresStudent = to.matched.some(
    (record) => record.meta.requiresStudent
  );
  const requiresTeacher = to.matched.some(
    (record) => record.meta.requiresTeacher
  );

  const isLogin = await store.getLoginState();
  if (requiresAuth && !isLogin) {
    alert("ログインしてください");
    next({ path: "/" });
    return;
  }

  next();
  return;
});

export default router;
