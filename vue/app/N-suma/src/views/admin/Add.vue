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
                  type="submit"
                  color="green"
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
                      color="red"
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
                      color="green"
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

const roles = ["学生", "担任"];
const clases = [
  "CTA20",
  "CTB20",
  "CTA21",
  "CTB21",
  "CTA22",
  "CTB22",
  "CTA23",
  "CTB23",
];

const requiredRules = [(v) => !!v || "必須"];

const numberRules = [
  (v) => !!v || "必須",
  (v) => v.length == 8 || "学籍番号は8桁です",
  (v) => /^\d+$/.test(v) || "学籍番号は半角数字です",
];

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
const uuids = ["1", "2", "3", "4"];
const emails = [
  "11@example.com",
  "22@example.com",
  "33@example.com",
  "44@example.com",
];

function emailSelected(e) {
  state.value.uuid = uuids[emails.indexOf(e)];
}

async function onSubmit() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // ユーザーを登録する処理を記述する
  console.log("ユーザーを登録しました");
}

onMounted(() => {
  console.log("mounted");
  // TODO: 役職が未定義のユーザー一覧を取得する(uuidとemailの組)
});
</script>
