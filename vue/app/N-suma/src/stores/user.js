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
import {
  onAuthStateChanged,
  signInWithRedirect,
  signOut,
} from "firebase/auth";

export const useStore = defineStore("user", () => {
  const user = ref(null);
  const role = ref(null);
  const isLogin = ref(false);

  onAuthStateChanged(auth, (u) => {
    // 取得したユーザー情報を格納する
    // ユーザーがログインしていない場合はnullを返す
    user.value = u;

    // ユーザー情報が取得できた場合に、役職を取得するAPIを呼び出す
    //ここで取得した役職をstoreに格納し、ログイン状態を管理する
    if (u) {
      // TODO: getRole()を実装する
      // role.value = getRole(u)
      role.value = "admin";

      isLogin.value = true;
    } else {
      role.value = null;
      isLogin.value = false;
    }
  });

  async function login() {
    await signInWithRedirect(auth, provider);
  }

  async function logout() {
    await signOut(auth);
  }

  return {
    user,
    role,
    isLogin,
    login,
    logout,
  };
});
