import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      useruuid: null,
      userId: null,
      ClassId: null,
      DocId: null,
    };
  },
  mutations: {
    setUseruuid(state, useruuid) {
      state.useruuid = useruuid;
    },
    setUserId(state, userId) {
      state.userId = userId;
    },
    setClassId(state, ClassId) {
      state.ClassId = ClassId;
    },
    setDocId(state, DocId) {
      state.DocId = DocId;
    },
  },
  actions: {
    updateUseruuid({ commit }, useruuid) {
      commit("setUseruuid", useruuid);
    },
    updateUserId({ commit }, userId) {
      // ユーザIDをVuexのstateに保存する
      commit("setUserId", userId);
    },
    updateClassId({ commit }, ClassId) {
      commit("setClassId", ClassId);
    },
    updateDocId({ commit }, DocId) {
      commit("setDocId", DocId);
    },
  },
  getters: {
    getUseruuid(state) {
      return state.useruuid;
    },
    getUserId(state) {
      return state.userId;
    },
    getClassId(state) {
      return state.ClassId;
    },
    getDocId(state) {
      return state.DocId;
    },
  },
});

export default store;
