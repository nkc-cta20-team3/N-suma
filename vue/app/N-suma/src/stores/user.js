/**
 * store/index.js
 * included in `plugins/index.js`
 *
 * used pinia
 * documentation: https://pinia.vuejs.org/
 */
import { defineStore } from "pinia";
import { auth, provider } from "@/plugins/firebase";
import { onAuthStateChanged, signInWithRedirect, signOut } from "firebase/auth";
import router from "@/router";

export const useStore = defineStore("user", {
  state: () => {
    const id = null; // ユーザーID
    const user = null; // ユーザー情報
    const role = null; // 役職
    const isLogin = false; // ログイン状態
    const token = null; // JWT token
    return {
      id,
      user,
      role,
      isLogin,
      token,
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
      this.resetData();
      router.push("/");
    },
    // ログイン状態を確認し、ログインしていればユーザー情報を取得する
    getLoginState() {
      return new Promise((resolve, reject) => {
        const unsuscribe = onAuthStateChanged(
          auth,
          (u) => {
            if (u) {
              this.isLogin = true;
              this.user = u;

              // TODO: getIdAndRole()を実装する
              // getIdAndRole(u.uid)

              //this.role = "admin";
              //this.role = "student";
              this.role = "teacher";

              // TODO: getIDを実装する
              this.id = 4;

              // tokenを取得する
              u.getIdToken().then((token) => {
                this.token = token;
                // console.log("token is " + token);
              });

              resolve(true);
            } else {
              this.user = null;
              //console.log("user is none data");
            }
            resolve(false);
          },
          (e) => reject(e)
        );
        unsuscribe();
      });
    },
    resetData() {
      this.user = null;
      this.role = null;
      this.isLogin = false;
    },
  },
});
