/**
 * router/index.js
 * included in `plugins/index.js`
 */

import { createRouter, createWebHistory } from "vue-router";
import { useStore } from "@/stores/user";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: () => import("@/views/index.vue") },
    /*
    {
      path: "/document_list",
      component: () => import("@/views/document_list.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/document_form",
      component: () => import("@/views/document_form.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/document_auth",
      component: () => import("@/views/document_auth.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/document_accept",
      component: () => import("@/views/document_accept.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/admin_add",
      component: () => import("@/views/admin_add.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/admin_edit",
      component: () => import("@/views/admin_edit.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/admin_update",
      component: () => import("@/views/admin_update.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/practice",
      component: () => import("@/views/practice.vue"),
      meta: { requiresAuth: true },
    },*/
  ],
});

router.beforeEach(async (to, from, next) => {
  const store = useStore();

  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // TODO: storeからユーザー情報を取得する
    if (store.isLogin) {
      /*
      if (to.meta.requiresAdmin && userId !== "admin") {
        alert("管理者権限が必要です");
        next({ path: "/" });
      } else if (to.meta.requiresStudent && userId !== "student") {
        alert("学生権限が必要です");
        next({ path: "/" });
      } else if (to.meta.requiresTeacher && userId !== "teacher") {
        alert("教員権限が必要です");
        next({ path: "/" });
      } else {
        next();
      }
      */
      next();
    } else {
      alert("ログインし直してください");
      next({ path: "/" });
    }
  } else {
    next();
  }
});

export default router;
