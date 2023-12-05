<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 提出書類のデータ表示 -->
        <RowCard title="申請日" :text="state.date" />
        <RowCard title="公欠区分" text="公欠区分" />
        <RowCard title="申請時間" text="申請時間" />
        <RowCard title="場所" text="場所" />
        <RowCard
          title="学生コメント"
          text="確認よろしくお願いいたします。aaaaaaaaaaaaaaaaaaaaaaaaaa"
        />

        <!-- 教員側入力フォーム -->
        <v-form ref="mainForm">
          <!-- 必要欠席時間 -->
          <v-text-field
            v-model="state.absence"
            label="必要欠席時間"
            persistent-hint
            placeholder=""
            persistent-placeholder
            :rules="requiredRules"
            class="mt-4"
          ></v-text-field>

          <!-- 教員コメント入力 -->
          <!--※エラー・必要欠席時間とコメントが同期している-->
          <v-text-field
            v-model="state.comment"
            label="教員コメント"
            persistent-hint
            placeholder="確認しました。"
            persistent-placeholder
            :rules="requiredRules"
            :counter="200"
          ></v-text-field>

          <template class="d-flex flex-row justify-end text-black">
            <!-- 認可ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props" color="success"
                  >認可</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>認可しますか？</div>
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
                      >認可</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </template>
            </v-dialog>
            <!-- 却下ボタン -->
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn height="40" width="150" v-bind="props" color="warning"
                  >却下</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>却下しますか？</div>
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
                      >却下</v-btn
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
import RowCard from "@/components/StudentViewRowCard.vue";
import { requiredRules } from "@/utils";
import router from "@/router";

const mainForm = ref(null);
const state = ref({
  id: "",
  date: "",

  absence: "",
  comment: "",
});

async function onSubmit() {
  // 入力チェック
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // TODO: データを送信する処理を記述する
  console.log("提出しました");
}

onMounted(() => {
  console.log("mounted");

  state.value.id = router.currentRoute.value.params.id;
  console.log(state.value.id);

  state.value.date = state.value.id;

  // TODO: データを取得する処理を記述する
});
</script>
