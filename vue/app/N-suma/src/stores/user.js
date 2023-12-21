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
import { APICallonJWT } from "@/utils";

export const useStore = defineStore("user", {
  state: () => {
    const role = null; // 役職
    const isLogin = false; // ログイン状態
    const token = null; // JWT token
    return {
      role,
      isLogin,
      token,
    };
  },
  getters: {
    // 役職を取得する
    getRole: (state) => state.role,
    // ログイン状態を取得する
    getIsLogin: (state) => state.isLogin,
    // JWT tokenを取得する
    getToken: (state) => state.token,
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

              // tokenを取得する
              u.getIdToken().then((token) => {
                this.token = token;
                // console.log("token is " + token);

                // 役職名を取得する
                APICallonJWT("auth", {}).then((res) => {
                  // console.log(res);

                  // documentはnullになることがある
                  if (res.document) {
                    // console.log(res.document.post_name);
                    if (res.document.post_name === "管理者")
                      this.role = "admin";
                    else if (res.document.post_name === "教員")
                      this.role = "teacher";
                    else if (res.document.post_name === "学生")
                      this.role = "student";
                  } else {
                    this.role = null;
                    // console.log("role is null");
                  }

                  // console.log(this.role);
                });
              });

              resolve(true);
            }
            resolve(false);
          },
          (e) => reject(e)
        );
        unsuscribe();
      });
    },
    resetData() {
      this.role = null;
      this.isLogin = false;
      this.token = null;
    },
    changeRole() {
      // 役職名を取得するコードをコメントアウトしないと動作しない

      //console.log(this.role);
      if (this.role === "admin") {
        this.role = "student";
        alert("学生に切り替えました");
      } else if (this.role === "student") {
        this.role = "teacher";
        alert("教員に切り替えました");
      } else if (this.role === "teacher") {
        this.role = "admin";
        alert("管理者に切り替えました");
      }
      //console.log(this.role);
    },
  },
});
