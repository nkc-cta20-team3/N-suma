import { createApp } from "vue";
import App from "./App.vue";

import router from "./router";
import store from "./store";

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

// Vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const vuetify = createVuetify({
  components,
  directives,
});

createApp(App).use(router).use(store).use(vuetify).mount("#app");
