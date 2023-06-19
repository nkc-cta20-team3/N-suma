import { createStore } from 'vuex'

const store = createStore({
  state() {
    return {
      userId: null,
      ClassId: null
    }
  },
  mutations: {
    setUserId(state, userId) {
      state.userId = userId
    },
    setClassId(state, ClassId) {
      state.ClassId = ClassId
    }
  },
  actions: {
    updateUserId({ commit }, userId) {
      // ユーザIDをVuexのstateに保存する
      commit('setUserId', userId)
    }
  },
  getters: {
    getUserId(state) {
      return state.userId
    },
    getClassId(state) {
      return state.ClassId
    }
  }
})

export default store