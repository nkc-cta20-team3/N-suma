<template>
  <v-app id="inspire">
      <!-- ヘッダーコンポーネントを読み込む -->
      <Header />

    <!-- メイン -->
    <v-main>
      <v-row justify="center">
        <v-col cols="12" sm="10" md="8" lg="6">
          <v-card ref="form">
            <!-- テキストボックス -->
            <v-card-text>
              <v-text-field
                ref="name"
                v-model="name"
                :rules="[() => !!name || '必須']"
                :error-messages="errorMessages"
                label="氏名"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="student_number"
                v-model="student_number"
                :rules="[() => !!student_number || '必須']"
                :error-messages="errorMessages"
                label="学籍番号"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="indispensable"
                v-model="indispensable"
                :rules="[() => !!indispensable || '必須']"
                :error-messages="errorMessages"
                label="クラス略称"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="section"
                v-model="position"
                :rules="[() => !!position || '必須']"
                :error-messages="errorMessages"
                label="役職"
                placeholder=""
                required
              ></v-text-field>
            </v-card-text>
            <v-divider class="mt-12"></v-divider>

            <!-- ボタン処理 -->
            <v-card-actions>
              <!-- 右寄せ用 -->
              <v-spacer></v-spacer>
              
              <!-- キャンセルボタン -->
              <v-btn variant="text"> Cancel </v-btn>
              <v-slide-x-reverse-transition>
                <v-tooltip v-if="formHasErrors" location="left">
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn
                      icon
                      class="my-0"
                      v-bind="attrs"
                      @click="resetForm"
                      v-on="on"
                    >
                      <v-icon>mdi-refresh</v-icon>
                    </v-btn>
                  </template>
                  <span>Refresh form</span>
                </v-tooltip>
              </v-slide-x-reverse-transition>
              
              <!-- 登録ボタン -->
              <v-container fluid>
                <v-btn color="green" @click="showAlertDialog">登録</v-btn>
                <v-dialog v-model="dialog" max-width="400">
                  <v-card>
                    <v-card-title>ユーザーを登録しますか？</v-card-title>
                    <v-card-actions>
                      <v-btn @click="registerUser">登録</v-btn>
                      <v-btn @click="cancelRegistration">キャンセル</v-btn>
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-container>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-main>
  </v-app>


</template>

<script>
//ヘッダー読み込み用
import Header from '../components/Navigation.vue';

export default {
  data: () => ({
    components:{
      Header,
    }}),

  //アラートボックス処理
  data() {
    return {
      dialog: false,
      name: '',
      student_number: '',
      indispensable: '',
      position: '',
      errorMessages: [],
    };
  },
  methods: {
    showAlertDialog() {
      this.dialog = true;
    },
    registerUser() {
      // ユーザーを登録する処理を記述する
      console.log('ユーザーを登録しました');
      this.dialog = false; // ダイアログを閉じる
    },
    cancelRegistration() {
      // キャンセル処理を記述する
      console.log('登録をキャンセルしました');
      this.dialog = false; // ダイアログを閉じる
    }
  }
};
</script>
<!-- script setup>
// import { onMounted, ref } from "vue";
// import {
//   getAuth,
//   onAuthStateChanged,
//   signOut,
//   GoogleAuthProvider,
//   signInWithRedirect,
// } from "firebase/auth";
// import router from "../router";
// //import store from './store';

// let auth;
// onMounted(() => {
//   auth = getAuth();
//   onAuthStateChanged(auth, (user) => {
//     isLoggedIn.value = !!user;

//     if (user) {
//       store.dispatch("updateUser", {
//         userId: userId,
//         classId: ClassId,
//         DocId: DocId,
//       });
//     }
//   });
// });
</script -->
