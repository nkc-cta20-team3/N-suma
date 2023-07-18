import { createStore } from 'vuex'

const store = createStore({
  state() {
    return {
      userId: null,
      ClassId: null,
      DocId: null,
      PostId: null
    }
  },
  mutations: {
    setUserId(state, userId) {
      state.userId = userId
    },
    setClassId(state, ClassId) {
      state.ClassId = ClassId
    },
    setDocId(state, DocId) {
      state.DocId = DocId
    },
    setPostId(state, PostId){
      state.PostId = PostId
    }
  },
  actions: {
    updateUserId({ commit }, userId) {
      // ユーザIDをVuexのstateに保存する
      commit('setUserId', userId)
    },
    updateClassId({ commit }, ClassId) {
      commit('setClassId', ClassId)
    },
    updateDocId({ commit }, DocId) {
      commit('setDocId', DocId)
    },
    updatePostId({ commit }, PostId) { 
      commit('setPostId', PostId)
    }
  },
  getters: {
    getUserId(state) {
      return state.userId
    },
    getClassId(state) {
      return state.ClassId
    },
    getDocId(state) {
      return state.DocId
    },
    getPostId(state){
      return state.PostId
    }
  }
})

export default store