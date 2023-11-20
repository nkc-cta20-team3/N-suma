<template>
  <v-app-bar color="primary">
    <!-- ロゴ -->
    <v-img
      class="bg-orange-lighten-4"
      src="/logo.svg"
      max-width="112"
      max-height="40"
    ></v-img>

    <v-btn v-if="isLoggedIn" @click="consoleDebug">debug</v-btn>

    <!-- メニュー項目 -->
    <!--
    <v-toolbar-items>
      <v-btn v-if="!isLoggedIn" to="/">ホーム</v-btn>

      <v-btn v-if="isLoggedIn" @click="consoleDebug">debug</v-btn>

      <v-btn v-if="isLoggedIn && userId !== 'admin'" to="/document_form"
        >各種書類提出</v-btn
      >

      <v-btn v-if="isLoggedIn && userId === 'student'" to="/document_list"
        >各種書類閲覧
      </v-btn>
      <v-btn v-if="isLoggedIn && userId === 'teacher'" to="/document_auth"
        >書類認可
      </v-btn>
      <v-btn v-if="isLoggedIn && userId === 'admin'" to="/admin_add"
        >ユーザー登録
      </v-btn>
      <v-btn v-if="isLoggedIn && userId === 'admin'" to="/admin_edit"
        >ユーザー情報編集
      </v-btn>
      <v-btn v-if="!isLoggedIn" @click="signInWithGoogle">ログイン </v-btn>
      <v-btn v-if="isLoggedIn" @click="handleSignOut">ログアウト</v-btn>
    </v-toolbar-items>
    -->
  </v-app-bar>
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

const isLoggedIn = ref(false);

let auth;
onMounted(() => {
  auth = getAuth();
  onAuthStateChanged(auth, (user) => {
    isLoggedIn.value = !!user;
  });
});

const handleSignOut = () => {
  signOut(auth)
    .then(() => {
      router.push("/");
      isLoggedIn.value = false;
    })
    .catch((error) => {
      console.log(error);
    });
};

const signInWithGoogle = () => {
  const provider = new GoogleAuthProvider();
  signInWithRedirect(getAuth(), provider)
    .then((result) => {
      //console.log(result.user)
      router.push("/document_form");
    })
    .catch((error) => {
      console.log(error);
    });
};

const consoleDebug = () => {
  console.log("store user state");
  console.log("====================");

  console.log("isLoggedIn ");
  console.log("====================");
  console.log(isLoggedIn.value);
  console.log("====================");

  console.log("auth ");
  console.log("====================");
  console.log(auth);
  console.log("====================");

  console.log("userId ");
  console.log("====================");
  console.log(userId);
  console.log("====================");
};
</script>
