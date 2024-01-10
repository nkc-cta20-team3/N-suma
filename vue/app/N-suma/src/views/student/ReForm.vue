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
import { requiredRules, APICallonGET, APICallonJWT } from "@/utils";

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

// 区分一覧を格納する変数
var divisions = ref([]);
var divisionids = [];

async function onSubmit() {
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

  // undifinedになるので、再度idを取得する
  // TODO: なぜundifinedになるのか調査する
  let id = router.currentRoute.value.params.id;
  console.log(id);

  // データを送信する処理を記述する
  console.log(state.value);
  APICallonJWT("student/reform/update", {
    document_id: Number(id),
    start_time: state.value.startDate,
    end_time: state.value.endDate,
    location: state.value.location,
    student_comment: state.value.studentComment,
    division_id: divisionids[divisions.value.indexOf(state.value.division)],
  }).then((res) => {
    console.log(res);
    if (res.message == "success") {
      alert("書類を再提出しました");
      router.push("/student/list");
    } else {
      alert("書類の再提出に失敗しました");
    }
  });
}

onMounted(() => {
  // console.log("mounted");
  state.value.id = router.currentRoute.value.params.id;
  init();
});

onBeforeRouteUpdate((to, from, next) => {
  // console.log("beforeRouteUpdate");
  state.value.id = to.params.id;
  init();
  next();
});

function init() {
  // 区分一覧を取得する
  APICallonGET("utils/read/division").then((res) => {
    // console.log(res);
    res.document.forEach((division_) => {
      divisions.value.push(
        division_.division_name + "/" + division_.division_detail
      );
      divisionids.push(division_.division_id);
    });
  });

  // 書類の詳細情報を取得する
  APICallonJWT("student/reform/read", {
    document_id: Number(state.value.id),
  }).then((res) => {
    // console.log(res);
    state.value = {
      division: res.document.division_name + "/" + res.document.division_detail,
      location: res.document.location,
      studentComment: res.document.student_comment,
      teacherComment: res.document.teacher_comment,
    };
  });
}
</script>
