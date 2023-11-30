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
      <v-btn
        @click.stop="dialog = true"
        icon
        v-if="store.role === 'student' || store.role === 'teacher'"
      >
        <!-- icon url : https://pictogrammers.com/library/mdi/icon/bell-outline/ -->
        <v-icon alt=" icon" :icon="mdiBellOutline"></v-icon>
      </v-btn>

      <!-- Login Button -->
      <v-btn v-if="!store.isLogin" @click="store.login">ログイン</v-btn>
      <v-btn v-if="store.isLogin" @click="store.logout">ログアウト</v-btn>
    </v-toolbar-items>
  </v-app-bar>
  <!-- 通知ポップアップ -->
  <v-dialog width="70%" v-model="dialog" scrollable>
    <v-card>
      <v-card-title v-if="!isNotification">通知はありません</v-card-title>
      <v-container v-else>
        <v-row justify="center">
          <v-col cols="12" v-for="item in notification" :key="item.title">
            <RowCard :id="item.id" :title="item.title" :text="item.text" />
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { mdiBellOutline } from "@mdi/js";
import RowCard from "@/components/NavigationRowCard.vue";
import { useStore } from "@/stores/user";
const store = useStore();

const dialog = ref(false);
const isNotification = ref(false);
const notification = ref([]);

function consoleDebug() {
  console.log("debug logs");
  console.log("====================");
  console.log("isLogin: ", store.isLogin);
  console.log("role: ", store.role);
  console.log("user: ", store.user ? store.user.uid : "null");
  console.log("====================\n");
}

onMounted(() => {
  // TODO: 通知がないかを確認するAPIを叩く処理を記述する
  isNotification.value = true;

  // MEMO: 叩くAPIは、学生か教員かで挙動が変わる
  // TODO: 通知がある場合は、notificationに通知の内容を格納する
  if (isNotification.value) {
    notification.value = [
      {
        id: 1,
        title: "通知1",
        text: "通知1の内容",
      },
      {
        id: 2,
        title: "通知2",
        text: "通知2の内容",
      },
      {
        id: 3,
        title: "通知3",
        text: "通知3の内容",
      },
    ];
  }
});
</script>
