<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 入力フォーム -->
        <v-form ref="mainForm">
          <!-- UUID選択 -->
          <v-text-field
            v-model="state.uuid"
            label="uuid"
            persistent-hint
            :disabled="true"
          ></v-text-field>

          <!-- Eメール選択 -->
          <v-text-field
            v-model="state.email"
            label="Eメール"
            persistent-hint
            :disabled="true"
          ></v-text-field>

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
            placeholder="学生"
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
          ></v-text-field>

          <!-- 所属クラス選択 -->
          <v-select
            v-model="state.class"
            :items="clases"
            label="所属クラス"
            persistent-hint
            placeholder="CTA20"
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

          <!-- アカウントの状態 -->
          <v-select
            v-model="state.status"
            :items="status"
            label="アカウントの状態"
            persistent-hint
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

          <!-- 各種ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <!-- 削除ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn
                  height="40"
                  width="150"
                  v-bind="props"
                  type="submit"
                  color="warning"
                  class="mr-2"
                  >削除</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>ユーザーを削除しますか?</div>
                  </v-card-text>
                  <v-card-actions class="justify-center">
                    <v-btn
                      height="40"
                      width="150"
                      @click="isActive.value = false"
                      color="success"
                      >キャンセル</v-btn
                    >
                    <v-btn
                      height="40"
                      width="150"
                      @click="
                        () => {
                          isActive.value = false;
                          onDelete();
                        }
                      "
                      color="warning"
                      >削除</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </template>
            </v-dialog>
            <!-- 更新ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn
                  height="40"
                  width="150"
                  v-bind="props"
                  type="submit"
                  color="success"
                  >更新</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>ユーザーを更新しますか?</div>
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
                          onUpdate();
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
  uuid: "",
  email: "",
  name: "",
  number: "",
  class: "",
  role: "",
  status: "",
});
var status = ["有効", "無効"];

// クラス一覧を格納する変数
var clases = ref([]);
var claseids = [];

// 役職一覧を格納する変数
var roles = ref([]);
var roleids = [];

async function onUpdate() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // TODO:ユーザーを更新する処理を記述する
  console.log("ユーザーを更新しました");
}

async function onDelete() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // TODO:ユーザーを削除する処理を記述する
  console.log("ユーザーを削除しました");
}

onMounted(() => {
  init();
});

onBeforeRouteUpdate((to, from, next) => {
  init();
  next();
});

function init() {
  // routerのpropsからidを取得する
  state.value.id = router.currentRoute.value.params.id;

  // クラス一覧を取得する
  APICallonGET("utils/read/class").then((res) => {
    // console.log(res);
    res.document.forEach((class_) => {
      clases.value.push(class_.class_name);
      claseids.push(class_.class_id);
    });
  });

  // ロール一覧を取得する
  APICallonGET("utils/read/post").then((res) => {
    // console.log(res);
    res.document.forEach((post) => {
      roles.value.push(post.post_name);
      roleids.push(post.post_id);
    });
  });

  // TODO: ユーザーの詳細な情報を取得し設定する
}
</script>
