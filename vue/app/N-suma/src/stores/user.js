/**
 * store/index.js
 * included in `plugins/index.js`
 *
 * used pinia
 * documentation: https://pinia.vuejs.org/
 */
import { ref } from "vue";
import { defineStore } from "pinia";
import { auth, provider } from "@/plugins/firebase";
import { onAuthStateChanged, signInWithRedirect, signOut } from "firebase/auth";

export const useStore = defineStore("user", {
  state: () => {
    const user = null;
    const role = null;
    const isLogin = false;

    return {
      user,
      role,
      isLogin,
    };
  },
  getters: {
    // ユーザー情報を取得する
    getUser: (state) => state.user,
    // 役職を取得する
    getRole: (state) => state.role,
    // ログイン状態を取得する
    getIsLogin: (state) => state.isLogin,
  },
  actions: {
    // ログインする
    async login() {
      await signInWithRedirect(auth, provider);
    },
    // ログアウトする
    async logout() {
      await signOut(auth);
    },
  },
});
