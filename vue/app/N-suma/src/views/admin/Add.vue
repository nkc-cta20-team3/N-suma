<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 入力フォーム -->
        <v-form ref="mainForm" v-if="emails.length > 0">
          <!-- UUID選択 -->
          <v-select
            v-model="state.uuid"
            label="uuid"
            persistent-hint
            :disabled="true"
          ></v-select>

          <!-- Eメール選択 -->
          <v-select
            v-model="state.email"
            :items="emails"
            label="Eメール"
            persistent-hint
            placeholder="aaaa1234@example.com"
            persistent-placeholder
            :rules="requiredRules"
            @update:modelValue="emailSelected"
          ></v-select>

          <!-- 氏名入力 -->
          <v-text-field
            v-model="state.name"
            label="氏名"
            persistent-hint
            placeholder="田中太郎"
            persistent-placeholder
            :rules="requiredRules"
          ></v-text-field>

          <!-- 役職選択 -->
          <v-select
            v-model="state.role"
            :items="roles"
            label="役職"
            persistent-hint
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

          <!-- 学籍番号入力 -->
          <v-text-field
            v-if="state.role == '学生'"
            v-model="state.number"
            label="学籍番号"
            persistent-hint
            placeholder="20201001"
            persistent-placeholder
            :rules="numberRules"
            :counter="8"
          ></v-text-field>

          <!-- 所属クラス選択 -->
          <v-select
            v-model="state.class"
            :items="clases"
            label="所属クラス"
            persistent-hint
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

          <!-- 登録ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props" color="success"
                  >登録</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>ユーザーを登録しますか?</div>
                  </v-card-text>
                  <v-card-actions class="justify-center">
                    <v-btn
                      height="40"
                      width="150"
                      @click="isActive.value = false"
                      color="warning"
                      >キャンセル</v-btn
                    >
                    <v-btn
                      height="40"
                      width="150"
                      @click="
                        () => {
                          isActive.value = false;
                          onSubmit();
                        }
                      "
                      color="success"
                      >登録</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </template>
            </v-dialog>
          </template>
        </v-form>
        <p v-else class="text-h5 text-center pt-10">
          新規登録対象のユーザーは存在しません。
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { onBeforeRouteUpdate } from "vue-router";
import router from "@/router";
import {
  requiredRules,
  numberRules,
  APICallonJWT,
  APICallonGET,
} from "@/utils";

const mainForm = ref(null);
const state = ref({
  id: "",
  name: "",
  number: "",
  role: "",
  class: "",
  email: "",
});

// ユーザー一覧を格納する変数
var userids = [];
var uuids = [];
var emails = ref([]);

// クラス一覧を格納する変数
var clases = [];
var claseids = [];

// 役職一覧を格納する変数
var roles = [];
var roleids = [];

function emailSelected(e) {
  var index = emails.value.indexOf(e);
  state.value.id = userids[index];
  state.value.uuid = uuids[index];
}

async function onSubmit() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // ユーザーを登録する処理
  APICallonJWT("admin/add/create", {
    user_id: state.value.id,
    user_name: state.value.name,
    user_number: state.role.value != "学生" ? null : state.value.number,
    post_id: roleids[roles.indexOf(state.value.role)],
    class_id: claseids[clases.indexOf(state.value.class)],
  }).then((res) => {
    res.message == "success"
      ? alert("ユーザーを登録しました")
      : alert("ユーザーの登録に失敗しました");
    // 画面遷移
    router.push("/app/admin/list");
  });
}

onMounted(() => {
  init();
});

onBeforeRouteUpdate((to, from, next) => {
  init();
  next();
});

function init() {
  // 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
  APICallonJWT("admin/add/read", {}).then((res) => {
    // console.log(res);
    res.document.forEach((user) => {
      userids.push(user.user_id);
      uuids.push(user.user_uuid);
      emails.value.push(user.mail_address);
    });
  });

  // クラス一覧を取得する
  APICallonGET("utils/read/class").then((res) => {
    // console.log(res);
    res.document.forEach((class_) => {
      clases.push(class_.class_name);
      claseids.push(class_.class_id);
    });
  });

  // ロール一覧を取得する
  APICallonGET("utils/read/post").then((res) => {
    // console.log(res);
    res.document.forEach((post) => {
      roles.push(post.post_name);
      roleids.push(post.post_id);
    });
  });
}
</script>
