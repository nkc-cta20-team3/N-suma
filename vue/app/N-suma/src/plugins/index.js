/**
 * plugins/index.js
 * included in `./src/main.js`
 */

// Plugins
import vuetify from "@/plugins/vuetify";
import store from "@/stores";
import router from "@/router";

export function registerPlugins(app) {
  app.use(vuetify).use(store).use(router);
}
