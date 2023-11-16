<template>
  <v-app>
    <v-app-bar app color="#FF9800">
      <!-- ロゴ -->
      <v-img src="#" alt="ロゴ" max-width="112" max-height="40"></v-img>
      <v-img src="/vite.svg" max-width="112" max-height="28"></v-img>

      <!-- メニュー項目 -->
      <v-toolbar-items>
        <v-btn v-if="!isLoggedIn" to="/">ホーム</v-btn>
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
    </v-app-bar>
  </v-app>
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
const isLoggedIn = ref(false);
const userId = "admin";
const classId = "";
const DocId = "";

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

//ハンバーガーメニューの実装
document.addEventListener("DOMContentLoaded", () => {
  // トグルボタンを取得
  const $navbarBurgers = Array.prototype.slice.call(
    document.querySelectorAll(".navbar-burger"),
    0
  );

  // トグルボタンが存在する場合
  if ($navbarBurgers.length > 0) {
    // トグルボタンにクリックイベントを設定
    $navbarBurgers.forEach((el) => {
      el.addEventListener("click", () => {
        // data-targetの属性値からナビゲーションメニューを取得
        const target = el.dataset.target;
        const $target = document.getElementById(target);

        // トグルボタンとナビゲーションメニューにis-activeクラスを設定
        el.classList.toggle("is-active");
        $target.classList.toggle("is-active");
      });
    });
  }
});
</script>

<!-- 旧環境のやつ -->
<!-- <template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <router-link class="navbar-item" to="/">
        <img src="#" alt="logo" width="112" height="40" />

        <img src="/vite.svg" width="112" height="28" />
      </router-link>

      ハンバーガーメニュー
      <a
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="navbarBasicExample"
      >
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

     メニュー項目 
    <div id="navbarBasicExample" class="navbar-menu">
      <div class="navbar-start">
        <router-link class="navbar-item" to="/" v-if="!isLoggedIn">
          Home
        </router-link>

        <router-link
          class="navbar-item"
          to="/document_form"
          v-if="isLoggedIn && userId !== 'admin'"
        >
          各種書類提出
        </router-link>

        <router-link
          class="navbar-item"
          to="/document_list"
          v-if="isLoggedIn && userId === 'student'"
        >
          各種書類閲覧
        </router-link>

        <router-link
          class="navbar-item"
          to="/document_auth"
          v-if="isLoggedIn && userId === 'teacher'"
        >
          書類認可
        </router-link>

        <router-link
          class="navbar-item"
          to="admin_add"
          v-if="isLoggedIn && userId === 'admin'"
        >
          ユーザー登録
        </router-link>

        <router-link
          class="navbar-item"
          to="admin_edit"
          v-if="isLoggedIn && userId === 'admin'"
        >
          ユーザー情報編集
        </router-link>
      </div>

      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
            <a
              @click="signInWithGoogle"
              v-if="!isLoggedIn"
              class="button is-primary"
            >
              <strong>Log in</strong>
            </a>
            <span class="material-symbols-outlined" v-if="isLoggedIn">
              notifications
            </span>
            <a @click="handleSignOut" v-if="isLoggedIn" class="button is-light">
              <strong>Sign Out</strong>
            </a>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template> -->
