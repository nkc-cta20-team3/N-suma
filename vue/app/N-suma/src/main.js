// Plugins
import { registerPlugins } from "@/plugins";

import { createApp } from "vue";
import App from "@/App.vue";

const app = createApp(App);
registerPlugins(app);
app.mount("#app");


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
