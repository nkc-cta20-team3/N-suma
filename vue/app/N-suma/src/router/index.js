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
              name: "adminEdit",
              path: "edit/:id",
              component: () => import("@/views/admin/Edit.vue"),
              props: true,
            },
            {
              path: "",
              redirect: "/",
            },
            {
              path: "/:pathMatch(.*)*",
              redirect: "/",
            },
          ],
          meta: { requiresAdmin: true },
        },
        {
          path: "student",
          component: () => import("@/layouts/Main.vue"),
          children: [
            {
              path: "form",
              component: () => import("@/views/student/Form.vue"),
            },
            {
              path: "list",
              component: () => import("@/views/student/List.vue"),
            },
            {
              name: "studentView",
              path: "view/:id",
              component: () => import("@/views/student/View.vue"),
              props: true,
            },
            {
              name: "studentReForm",
              path: "reform/:id",
              component: () => import("@/views/student/ReForm.vue"),
              props: true,
            },
            {
              path: "",
              redirect: "/",
            },
            {
              path: "/:pathMatch(.*)*",
              redirect: "/",
            },
          ],
          meta: { requiresStudent: true },
        },
        {
          path: "teacher",
          component: () => import("@/layouts/Main.vue"),
          children: [
            {
              path: "unapproval",
              component: () => import("@/views/teacher/UnApprovalList.vue"),
            },
            {
              name: "auth",
              path: "auth/:id",
              component: () => import("@/views/teacher/auth.vue"),
              props: true,
            },
            {
              path: "list",
              component: () => import("@/views/teacher/List.vue"),
            },
            {
              name: "teacherView",
              path: "view/:id",
              component: () => import("@/views/teacher/View.vue"),
              props: true,
            },
            // 現状使用していない
            // auth.vueに該当するもの
            /*
            {
              
              name: "teacherForm",
              path: "form/:id",
              component: () => import("@/views/teacher/Form.vue"),
              props: true,
            },
            */
            {
              path: "",
              redirect: "/",
            },
            {
              path: "/:pathMatch(.*)*",
              redirect: "/",
            },
          ],
          meta: { requiresTeacher: true },
        },
        {
          path: "",
          redirect: "/",
        },
        {
          path: "/:pathMatch(.*)*",
          redirect: "/",
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

  await store.fetchRole();
  if (
    (requiresAdmin && store.role != "admin") ||
    (requiresStudent && store.role != "student") ||
    (requiresTeacher && store.role != "teacher")
  ) {
    alert("権限がありません");
    next({ path: "/" });
    return;
  }

  next();
  return;
});

export default router;
