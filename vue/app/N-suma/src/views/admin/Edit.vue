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
            :disabled="true"
          ></v-select>

          <!-- Eメール選択 -->
          <v-select
            v-model="state.email"
            label="Eメール"
            :disabled="true"
          ></v-select>

          <!-- 氏名入力 -->
          <v-text-field
            v-model="state.name"
            label="氏名"
            persistent-hint
            placeholder="田中太郎"
            persistent-placeholder
            :rules="requiredRule"
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

          <!-- TODO&MEMO:クラス略称の一覧が入手できた場合、ハードコーディングしてしまった方がよさそう -->
          <!-- クラス略称選択 -->
          <v-select
            v-model="state.class"
            :items="classes"
            label="クラス略称"
            persistent-hint
            placeholder="CTA20"
            persistent-placeholder
            :rules="requiredRule"
          ></v-select>

          <!-- 役職選択 -->
          <v-select
            v-model="state.role"
            :items="roles"
            label="役職"
            persistent-hint
            placeholder="学生"
            persistent-placeholder
            :rules="requiredRule"
          ></v-select>

          <!-- 各種ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <!-- 更新ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props">更新</v-btn>
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
                      >更新</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </template>
            </v-dialog>
            <!-- 削除ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props">削除</v-btn>
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
                      >削除</v-btn
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

const mainForm = ref(null);

const state = ref({
  uuid: "1",
  email: "11@example.com",
  name: "",
  number: "",
  class: "",
  role: "",
});

const classes = ["CTA20", "CTB20", "CTC20", "CTD20", "CTE20", "CTF20"];
const roles = ["担任", "学生"];

const requiredRule = [(v) => !!v || "必須"];

const numberRules = [
  (v) => v.length == 8 || "学籍番号は8桁です",
  (v) => /^\d+$/.test(v) || "学籍番号は半角数字です",
  (v) => !!v || "必須",
];

async function onUpdate() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("ユーザーを更新できませんでした");
  }
  // TODO:ユーザーを更新する処理を記述する
  console.log("ユーザーを更新しました");
}

async function onDelete() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("ユーザーを削除できませんでした");
    return;
  }
  // TODO:ユーザーを削除する処理を記述する
  console.log("ユーザーを削除しました");
}

onMounted(() => {
  // TODO:ユーザー情報(uuid&email)をpropsから取得する処理を記述する。存在しない場合、listに遷移する。
});
</script>
