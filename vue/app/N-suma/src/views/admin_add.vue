<template>
  <div>
    <h1>
      管理者アカウント登録画面だお　ぶいすぽっ文化体育祭行きたかった！！！！！！！！！
    </h1>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import {
  getAuth,
  onAuthStateChanged,
  signOut,
  GoogleAuthProvider,
  signInWithRedirect,
} from "firebase/auth";
import router from "../router";
//import store from './store';

let auth;
onMounted(() => {
  auth = getAuth();
  onAuthStateChanged(auth, (user) => {
    isLoggedIn.value = !!user;

    if (user) {
      store.dispatch("updateUser", {
        userId: userId,
        classId: ClassId,
        DocId: DocId,
      });
    }
  });
});
</script>
