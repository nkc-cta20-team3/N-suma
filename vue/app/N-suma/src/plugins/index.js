/**
 * plugins/index.js
 * included in `./src/main.js`
 */

// Plugins
import vuetify from './vuetify.js'
import router from '../router'
import store from "../store";

export function registerPlugins (app) {
  app
    .use(vuetify)
    .use(router)
    .use(store)
}

