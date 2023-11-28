<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 入力フォーム -->
        <v-form @submit.prevent="onSubmit">
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
            :rules="[(v) => !!v || '必須']"
            @update:modelValue="emailSelected"
          ></v-select>

          <!-- 氏名入力 -->
          <v-text-field
            v-model="state.name"
            label="氏名"
            persistent-hint
            placeholder="田中太郎"
            persistent-placeholder
            :rules="[(v) => !!v || '必須']"
          ></v-text-field>

          <!-- 学籍番号入力 -->
          <v-text-field
            v-model="state.number"
            label="学籍番号"
            persistent-hint
            placeholder="20201001"
            persistent-placeholder
            :rules="numberRules"
          ></v-text-field>

          <!-- クラス略称入力 -->
          <v-text-field
            v-model="state.class"
            label="クラス略称"
            persistent-hint
            placeholder="CTA20"
            persistent-placeholder
            :rules="classRules"
          ></v-text-field>

          <!-- 役職選択 -->
          <v-select
            v-model="state.role"
            :items="roles"
            label="役職"
            persistent-hint
            placeholder="学生"
            persistent-placeholder
            :rules="[(v) => !!v || '必須']"
          ></v-select>

          <!-- 登録ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props" type="submit"
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
                      >キャンセル</v-btn
                    >
                    <v-btn
                      height="40"
                      width="150"
                      @click="isActive.value = false"
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

const state = ref({
  uuid: "",
  email: "",
  name: "",
  number: "",
  class: "",
  role: "",
});

// TODO: 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
const uuids = ["1", "2", "3", "4"];
const emails = ["11", "22", "33", "44"];
const roles = ["担任", "学生"];

const numberRules = [
  (v) => !!v || "必須",
  (v) => v.length == 8 || "学籍番号は8桁です",
  (v) => /^\d+$/.test(v) || "学籍番号は半角数字です",
];

// TODO:クラス略称は5桁なのか確認する
// それか、クラス略称をDBから取得する？
const classRules = [
  (v) => !!v || "必須",
  (v) => v.length == 5 || "クラス略称は5桁です",
  (v) => /^[A-Z0-9]+$/.test(v) || "クラス略称は大文字半角英字と半角数字です",
];

function emailSelected(e) {
  state.value.uuid = uuids[emails.indexOf(e)];
}

function onSubmit() {
  // ユーザーを登録する処理を記述する
  console.log("ユーザーを登録しました");
}

onMounted(() => {
  console.log("mounted");
  // TODO: 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
});
</script>
