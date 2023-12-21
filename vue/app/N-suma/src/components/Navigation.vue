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
        <v-btn to="/app/teacher/unapproval">書類認可</v-btn>
        <v-btn to="/app/teacher/list">書類閲覧</v-btn>
      </template>

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

      <!-- Debug Button -->
      <!-- <v-btn @click="consoleDebug">debug</v-btn> -->

      <!-- Debug Button changeRole -->
      <!-- <v-btn v-if="store.isLogin" @click="store.changeRole">役職切替</v-btn> -->
    </v-toolbar-items>
  </v-app-bar>
  <!-- 通知ポップアップ -->
  <v-dialog width="70%" v-model="dialog" scrollable>
    <v-card>
      <v-container v-if="isNotification">
        <v-row justify="center">
          <v-col cols="12" v-for="item in notification" :key="item.title">
            <RowCard
              @click="onItemClick(item.id)"
              :title="item.title"
              :text="item.text"
            />
          </v-col>
        </v-row>
      </v-container>
      <v-container v-else>通知はありません</v-container>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { mdiBellOutline } from "@mdi/js";
import RowCard from "@/components/NavigationRowCard.vue";
import { onBeforeRouteUpdate } from "vue-router";
import router from "@/router";
import { useStore } from "@/stores/user";
import { APICallonJWT } from "@/utils";

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

function onItemClick(id) {
  dialog.value = false;
  // console.log(props.id);

  // 学生の場合
  if (store.role === "student") {
    router.push({
      name: "studentReForm",
      params: { id: id },
    });
  } else if (store.role === "teacher") {
    router.push({
      name: "teacherForm",
      params: { id: id },
    });
  }
}

onMounted(() => {
  init();
});

onBeforeRouteUpdate((to, from, next) => {
  init();
  next();
});

function init() {
  // MEMO: 叩くAPIは、学生か教員かで挙動が変わる
  if (store.role === "student") {
    // 学生向けの通知がないかを確認するAPIを叩く
    APICallonJWT("student/alarm/check", {}).then((res) => {
      // console.log(res);
      isNotification.value = res.document;
    });
  } else if (store.role === "teacher") {
    // 教員向けの通知がないかを確認するAPIを叩く
    APICallonJWT("teacher/alarm/check", {}).then((res) => {
      // console.log(res);
      isNotification.value = res.document;
    });
  }

  console.log(isNotification.value);

  // TODO: 通知がある場合は、notificationに通知の内容を格納する
  if (isNotification) {
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
}
</script>
