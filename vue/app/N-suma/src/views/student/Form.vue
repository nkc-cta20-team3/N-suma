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
                <v-btn height="40" width="150" v-bind="props" color="success"
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
import { onMounted, ref } from "vue";
import router from "@/router";
import VueDatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { requiredRules, APICallonJWT, APICallonGET } from "@/utils";
import { useStore } from "@/stores/user";

const store = useStore();
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

  // データを送信する処理を記述する
  // console.log(state.value);
  APICallonJWT("student/form/create", {
    user_id: store.id,
    start_time: state.value.startDate,
    end_time: state.value.endDate,
    location: state.value.location,
    student_comment: state.value.comment,
    division_id: divisionids[divisions.value.indexOf(state.value.division)],
  }).then((res) => {
    // console.log(res);
    if (res.message == "success") {
      alert("書類を提出しました");
      router.push("/student/list");
    } else {
      alert("書類の提出に失敗しました");
    }
  });
}

onMounted(() => {
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
});
</script>
