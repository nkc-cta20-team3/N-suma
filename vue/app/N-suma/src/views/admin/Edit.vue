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

  // ユーザーを更新する処理を記述する
  APICallonJWT("admin/add/update", {
    user_id: state.value.id,
    user_name: state.value.name,
    user_number: state.role.value != "学生" ? null : state.value.number,
    post_id: roleids[roles.indexOf(state.value.role)],
    class_id: claseids[clases.indexOf(state.value.class)],
    user_flag: state.value.status == "有効" ? true : false,
  }).then((res) => {
    res.message == "success"
      ? alert("ユーザー情報を更新しました")
      : alert("ユーザー情報の更新に失敗しました");
    // 画面遷移
    router.push("/app/admin/edit");
  });
}

async function onDelete() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // ユーザーを削除する処理
  APICallonJWT("admin/edit/delete", {
    user_id: state.value.id,
  }).then((res) => {
    res.message == "success"
      ? alert("ユーザーの凍結に成功しました")
      : alert("ユーザーの凍結に失敗しました");
    // 画面遷移
    router.push("/app/admin/edit");
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

  // ユーザーの詳細な情報を取得し設定する
  APICallonJWT("admin/edit/read", {
    user_id: state.value.id,
  }).then((res) => {
    // console.log(res);
    if (res.message == "success") {
      state.value.uuid = res.document.user_uuid;
      state.value.email = res.document.mail_address;
      state.value.name = res.document.user_name;
      state.value.number = res.document.user_number;
      state.value.class = clases.value[claseids.indexOf(res.document.class_id)];
      state.value.role = roles.value[roleids.indexOf(res.document.post_id)];
      state.value.status = res.document.user_flag ? "有効" : "無効";
    } else {
      console.log("データの取得に失敗しました");
    }
  });
}
</script>
