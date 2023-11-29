<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 入力フォーム -->
        <v-form ref="mainForm">
          <!-- 公欠区分選択 -->
          <v-select
            v-model="state.division"
            :items="divisions"
            label="公欠区分"
            persistent-hint
            placeholder="国家試験 / FE"
            persistent-placeholder
            :rules="requiredRules"
          ></v-select>

          <!-- 公欠開始/終了時刻選択 -->
          <v-card width="100%"> 公欠開始/終了時刻 </v-card>
          <template class="d-flex flex-row text-black">
            <VueDatePicker
              v-model="state.startDate"
              :preview-format="format"
              :format="format"
              :enable-time-picker="true"
              week-start="0"
              locale="jp"
              placeholder="日付を選択"
              range
            ></VueDatePicker>
          </template>

          <!-- 実施場所入力 -->
          <v-text-field
            v-model="state.location"
            label="実施場所"
            persistent-hint
            placeholder="愛知県名古屋市熱田区神宮 / 名古屋工学院専門学校"
            persistent-placeholder
            :rules="requiredRules"
            class="mt-4"
          ></v-text-field>

          <!-- 学生コメント入力 -->
          <v-text-field
            v-model="state.comment"
            label="学生コメント"
            persistent-hint
            placeholder="確認よろしくお願いいたします。"
            persistent-placeholder
            :rules="requiredRules"
            :counter="200"
          ></v-text-field>

          <!-- 提出ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn
                  height="40"
                  width="150"
                  v-bind="props"
                  type="submit"
                  color="success"
                  >提出</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>提出しますか？</div>
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
                      >提出</v-btn
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
import VueDatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";

import { onMounted, ref } from "vue";

const requiredRules = [(v) => !!v || "必須"];
const divisions = ["国家試験 / FE", "国家試験 / AP"];
const format = "yyyy-MM-dd HH:mm";

const mainForm = ref(null);
const state = ref({
  date: "",
  division: "",
  startDate: "",
  endDate: "",
  location: "",
  comment: "",
});

async function onSubmit() {
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

  var date = new Date();
  state.value.date =
    date.getFullYear() + "-" + (date.getMonth() + 1) + "-" + date.getDate();
});
</script>
