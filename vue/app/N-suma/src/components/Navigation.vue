<template>
  <v-app-bar color="primary">
    <!-- Logo -->
    <v-img
      class="bg-orange-lighten-4"
      src="/logo.svg"
      min-width="112"
      min-height="40"
      max-width="112"
      max-height="40"
    ></v-img>

    <!-- Menu -->
    <v-toolbar-items>
      <!-- Home Button -->
      <v-btn to="/">home</v-btn>

      <!-- Admin Button -->
      <template v-if="store.role === 'admin'">
        <v-btn to="/app/admin/add">ユーザー登録</v-btn>
        <v-btn to="/app/admin/list">ユーザー情報編集</v-btn>
      </template>

      <!-- Student Button -->
      <template v-if="store.role === 'student'">
        <v-btn to="/app/student/form">書類提出</v-btn>
        <v-btn to="/app/student/list">書類閲覧</v-btn>
      </template>

      <!-- Teacher Button -->
      <template v-if="store.role === 'teacher'">
        <!-- TODO:適切なルートを設定する -->
        <v-btn to="/app/teacher">teacher</v-btn>
        <v-btn to="/app/teacher">teacher</v-btn>
        <v-btn to="/app/teacher">teacher</v-btn>
      </template>

      <!-- Debug Button -->
      <!-- <v-btn @click="consoleDebug">debug</v-btn> -->

      <!-- 通知ボタン -->
      <v-btn @click="onClickBell" icon>
        <!-- icon url : https://pictogrammers.com/library/mdi/icon/bell-outline/ -->
        <v-icon alt=" icon" :icon="mdiBellOutline"></v-icon>
      </v-btn>

      <!-- Login Button -->
      <v-btn v-if="!store.isLogin" @click="store.login">ログイン</v-btn>
      <v-btn v-if="store.isLogin" @click="store.logout">ログアウト</v-btn>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script setup>
import { mdiBellOutline } from "@mdi/js";
import { useStore } from "@/stores/user";
import { onMounted } from "vue";
const store = useStore();

function consoleDebug() {
  console.log("debug logs");
  console.log("====================");
  console.log("isLogin: ", store.isLogin);
  console.log("role: ", store.role);
  console.log("user: ", store.user ? store.user.uid : "null");
  console.log("====================\n");
}

function onClickBell() {
  console.log("通知ボタンが押されました");
  //TODO: 通知ボタンが押された時の処理を記述する
  // ダイアログボックス形式で通知の一覧を取得するのが良さそう？
}

onMounted(() => {
  // TODO: 通知がないかを確認するAPIを叩く処理を記述する
});
</script>
