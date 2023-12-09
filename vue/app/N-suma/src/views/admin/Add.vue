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

          <!-- 学籍番号入力 -->
          <v-text-field
            v-model="state.number"
            label="学籍番号"
            persistent-hint
            placeholder="20201001"
            persistent-placeholder
            :rules="numberRules"
            :counter="8"
          ></v-text-field>

          <!-- クラス略称選択 -->
          <v-select
            v-model="state.class"
            :items="clases"
            label="クラス略称"
            persistent-hint
            placeholder="CTA20"
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

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
        <p v-else class="text-h5 text-center pt-10">新規登録対象のユーザーは存在しません。</p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";
import {
  roles,
  clases,
  requiredRules,
  numberRules,
  APICallonJWT,
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
    user_number: Number(state.value.number),
    post_id: state.value.role == "学生" ? 1 : 2,
    // TODO:実装する
    // user_class: state.value.class,
    class_id: 132,
  }).then((res) => {
    res.message == "success"
      ? console.log("ユーザーを登録しました")
      : console.log("ユーザーの登録に失敗しました");
  });
}

onMounted(() => {
  // 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
  APICallonJWT("admin/add/read", {}).then((res) => {
    // console.log(res);
    res.document.forEach((user) => {
      userids.push(user.user_id);
      uuids.push(user.user_uuid);
      emails.value.push(user.mail_address);
    });
  });
});
</script>
