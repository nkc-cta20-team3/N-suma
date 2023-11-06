import { createApp } from "vue";
import App from "./App.vue";

import router from "./router";
import store from "./store";

// Firebase
import { getAuth, onAuthStateChanged } from "firebase/auth";

// Firebaseのログイン状態を監視し、VuexストアにユーザーIDを保存する
onAuthStateChanged(getAuth(), (user) => {
  if (user) {
    store.dispatch("updateUserId", user.uid);
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
