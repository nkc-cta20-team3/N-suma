import { createApp } from "vue";
import App from "./App.vue";

// Firebase
import { getAuth, onAuthStateChanged } from "firebase/auth";

// Firebaseのログイン状態を監視し、VuexストアにユーザーIDを保存する
onAuthStateChanged(getAuth(), (user) => {
  if (user) {
    store.dispatch("updateUseruuid", user.uuid);

    //  ユーザーの役職を取得するAPIを呼び出す
    const userId = 'admin'; // 予備
    // 取得したユーザーの役職をVuexのstateに保存する
    store.dispatch("updateUserId", userId);
    
  }
});

// Plugins
import { registerPlugins } from './plugins'

const app = createApp(App)
registerPlugins(app)
app.mount("#app");
