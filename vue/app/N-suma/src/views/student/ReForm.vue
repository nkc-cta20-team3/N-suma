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
              v-model="date"
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
            v-model="state.studentComment"
            label="学生コメント"
            persistent-hint
            placeholder="確認よろしくお願いいたします。"
            persistent-placeholder
            :rules="requiredRules"
            :counter="200"
          ></v-text-field>

          <!-- 教員コメント表示 -->
          <v-text-field
            v-model="state.teacherComment"
            label="教員コメント"
            persistent-hint
            :disabled="true"
          ></v-text-field>

          <!-- 再提出ボタン -->
          <template class="d-flex flex-row justify-end text-black">
            <v-dialog transition="dialog-top-transition" width="auto">
              <template v-slot:activator="{ props }">
                <v-btn
                  height="40"
                  width="150"
                  v-bind="props"
                  color="success"
                  class="mr-2"
                  >再提出</v-btn
                >
              </template>
              <template v-slot:default="{ isActive }">
                <v-card>
                  <v-card-text>
                    <div>書類を再提出しますか?</div>
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
import { onMounted, ref } from "vue";
import VueDatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { onBeforeRouteUpdate } from "vue-router";
import router from "@/router";
import { requiredRules, divisions, APICall } from "@/utils";

const format = "yyyy-MM-dd HH:mm";

const mainForm = ref(null);
const state = ref({
  id: "",
  division: "",
  startDate: "",
  endDate: "",
  location: "",
  studentComment: "",
  teacherComment: "",
});
const date = ref(null);

async function onSubmit() {
  console.log(state.value);

  // 日付の入力チェック
  if (!date.value || !date.value[0] || !date.value[1]) {
    console.log("日付入力不足");
    // TODO: エラーを画面に表示する処理を記述する？
    
    return;
  }
  // 日付以外の入力チェック
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }
  state.value.startDate = date.value[0];
  state.value.endDate = date.value[1];

  // TODO:書類を提出する処理を記述する
  const createdocument_url = "/api/student/cd";
  APICall("POST", createdocument_url, state);
  
  console.log(json);
  console.log("提出しました");

  // 画面遷移
  router.push("/student/list");
}

onMounted(() => {
  console.log("mounted");
  // TODO: routerのpropsから書類のidを取得し、stateに設定する
  state.value.id = router.currentRoute.value.params.id;
  init();
});

// TODO: 画面遷移時に呼ばれる処理を記述する
const readalarm_url = "/api/student/al";
  APICall("POST", readalarm_url, state);
  
  console.log(json);
onBeforeRouteUpdate((to, from, next) => {
  console.log("beforeRouteUpdate");
  // TODO: 画面更新時に呼び出されるため、to.params.idをもとにstateを初期化する
  state.value.id = to.params.id;
  init();
  next();
});

function init() {
  // TODO: 書類のidをもとに、APIを叩いて詳細な情報を取得し設定する
  const resubmitdocument_url = "/api/student/rsd";
  APICall("POST", resubmitdocument_url, state);
  
  console.log(json);

  state.value.division = "国家試験";
  state.value.location = "愛知県名古屋市熱田区神宮";
  state.value.studentComment = "確認よろしくお願いいたします。";
  state.value.teacherComment = "確認しました。";
}
</script>
