<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 入力フォーム -->
        <v-form ref="mainForm">
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
                <v-btn
                  height="40"
                  width="150"
                  v-bind="props"
                  
                  color="success"
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
      </v-col>
    </v-row>
  </v-container>
</template>
<script setup>
import { onMounted, ref } from "vue";
import { roles,clases, requiredRules, numberRules, APICall } from "@/utils";

const mainForm = ref(null);
const state = ref({
  uuid: "",
  email: "",
  name: "",
  number: "",
  class: "",
  role: "",
});

// TODO: 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
var uuids = [];
var emails = [];

function emailSelected(e) {
  state.value.uuid = uuids[emails.indexOf(e)];
}

async function onSubmit() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }
  console.log(state);

  //idとnumberを数値に変換
  state.value.uuid = Number(state.value.uuid);
  state.value.number = parseInt(state.value.number);

  //学生か教員かでpost_idに入れる値
  const insert_post_id = roles.indexOf(state.value.role)+1;
  console.log(insert_post_id);

  //TODO: クラス_idを数値に変化

  //ユーザーを登録する処理
  APICall("POST", "http://localhost:8080/api/admin/cu", {
    user_id: state.value.uuid,
    user_name: state.value.name,
    user_number: state.value.number,
    post_id: insert_post_id,
    class_id: 117,
    mail_address: state.value.email,
  }).then((res) => {
    console.log(res);
  });
  
  console.log(json);
  
  console.log("ユーザーを登録しました");
}

onMounted(() => {
  console.log("mounted");
  // TODO: 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
  uuids = ["1", "2", "3", "4"];
  emails = [
    "11@example.com",
    "22@example.com",
    "33@example.com",
    "44@example.com",
  ];
});

</script>