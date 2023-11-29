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
            label="Eメール"
            persistent-hint
            :disabled="true"
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
import router from "@/router";

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
  console.log("mounted");
  // TODO: routerのpropsからユーザー情報を取得する(uuidとemailの組)

  state.value.uuid = router.currentRoute.value.params.id;
  state.value.email = "11@example.com";

  // TODO: ユーザーの詳細な情報を取得し設定する
});
</script>
