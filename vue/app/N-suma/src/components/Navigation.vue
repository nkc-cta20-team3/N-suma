<template>
  <v-app-bar color="primary">
    <!-- ロゴ -->
    <v-img
      class="bg-orange-lighten-4"
      src="/logo.svg"
      max-width="112"
      max-height="40"
    ></v-img>

    <!-- メニュー項目 -->
    <v-toolbar-items>
      <v-btn @click="consoleDebug">debug</v-btn>
      <v-btn v-if="!store.isLogin" @click="login">ログイン</v-btn>
      <v-btn v-if="store.isLogin" @click="logout">ログアウト</v-btn>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script setup>
import { useStore } from "@/stores/user";
const store = useStore();

const consoleDebug = () => {
  console.log("debug logs");
  console.log("====================");
  console.log("isLogin: ", store.isLogin);
  console.log("role: ", store.role);
  console.log("user: ", store.user.uid);
  console.log("====================\n");
};

const login = () => {
  store.login();
};

const logout = () => {
  store.logout();
};
</script>
