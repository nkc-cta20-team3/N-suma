/**
 * plugins/store/index.js
 * included in `../index.js`
 * 
 * documentation: https://vuex.vuejs.org/ja/api/#createstore
 */

import { createStore } from "vuex";

export default createStore({
  // 状態を定義する
  state: { 
    // ユーザーの役職
    userPos: null,
  },
  mutations: {
    // ユーザーの役職をセットする
    setUserPos(state, userPos) {
      state.userPos = userPos;
    },
  },
  actions: {
  
  },
  getters: {
    // ユーザーの役職を取得する
    getUserPos(state) {
      return state.userPos;
    },
  },
  modules: {
  
  },
  plugins: [
    // デバッグモードの時のみ有効にする
    // process.env.NODE_ENV !== "production" && require("vuex/dist/logger")(),
  ], 
  strict: true,
});
