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
              path: "add",
              component: () => import("@/views/admin/Add.vue"),
            },
            {
              path: "list",
              component: () => import("@/views/admin/List.vue"),
            },
            {
              path: "edit",
              component: () => import("@/views/admin/Edit.vue"),
              // props: true,
              // props: (route) => ({ id: route.query.id }),
            },
          ],
          meta: { requiresAdmin: true },
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
