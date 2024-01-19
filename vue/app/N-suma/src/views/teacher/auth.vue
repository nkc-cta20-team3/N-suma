<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 提出書類のデータ表示 -->
        <RowCard title="申請日" :text="state.date" />
        <RowCard title="公欠区分" :text="state.division" />
        <RowCard title="申請時間" :text="state.dateTime" />
        <RowCard title="場所" :text="state.location" />
        <RowCard title="学生コメント" :text="state.studentComment" />

        <!-- 教員側入力フォーム -->
        <v-form ref="mainForm">
          <!-- 必要欠席時間 -->
          <v-row gutter="0">
            <v-col cols="6">
              <v-text-field
                v-model="state.startAbsence"
                label="公欠開始時限"
                persistent-hint
                placeholder=""
                persistent-placeholder
                :rules="requiredRules"
              >
              </v-text-field>
            </v-col>
            <v-col cols="6">
              <v-text-field
                v-model="state.endAbsence"
                label="公欠終了時限"
                persistent-hint
                placeholder=""
                persistent-placeholder
                :rules="requiredRules"
              ></v-text-field>
            </v-col>
          </v-row>

          <!-- 教員コメント入力 -->
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
                          onAuth();
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
                          onReject();
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
import { onBeforeRouteUpdate } from "vue-router";
import RowCard from "@/components/StudentViewRowCard.vue";
import { requiredRules } from "@/utils";
import router from "@/router";
import { APICallonJWT } from "@/utils";

const mainForm = ref(null);
const state = ref({
  id: "",
  date: "",
  division: "",
  startDate: "",
  endDate: "",
  dateTime: "",

  location: "",
  studentComment: "",

  startAbsence: "",
  endAbsence: "",
  comment: "",
});

async function onAuth() {
  // 入力チェック
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // undifinedになるので、再度idを取得する
  // TODO: なぜundifinedになるのか調査する
  let id = router.currentRoute.value.params.id;

  // 書類を却下する処理
  APICallonJWT("teacher/auth/update", {
    document_id: Number(id),
    start_flame: Number(state.value.startAbsence),
    end_flame: Number(state.value.endAbsence),
    teacher_comment: state.value.comment,
  }).then((res) => {
    console.log(res);
    if (res.message == "success") {
      alert("書類を認可しました");
      router.push("/app/teacher/unapproval");
    } else {
      alert("書類の認可に失敗しました");
    }
  });
}

async function onReject() {
  // 入力チェック
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // undifinedになるので、再度idを取得する
  // TODO: なぜundifinedになるのか調査する
  let id = router.currentRoute.value.params.id;

  // 書類を却下する処理
  APICallonJWT("teacher/auth/reject", {
    document_id: Number(id),
    start_flame: Number(state.value.startAbsence),
    end_flame: Number(state.value.endAbsence),
    teacher_comment: state.value.comment,
  }).then((res) => {
    console.log(res);
    if (res.message == "success") {
      alert("書類を却下しました");
      router.push("/app/teacher/unapproval");
    } else {
      alert("書類の却下に失敗しました");
    }
  });
}

onMounted(() => {
  state.value.id = router.currentRoute.value.params.id;
  init();
});

onBeforeRouteUpdate((to, from, next) => {
  state.value.id = to.params.id;
  init();
  next();
});

function init() {
  // 書類の詳細情報を取得する
  APICallonJWT("teacher/auth/read", {
    document_id: Number(router.currentRoute.value.params.id),
  }).then((res) => {
    // console.log(res);
    state.value = {
      date: res.document.request_at,
      division: res.document.division_name + "/" + res.document.division_detail,
      startDate: res.document.start_time,
      endDate: res.document.end_time,
      location: res.document.location,
      studentComment: res.document.student_comment,
    };

    // 表示用に日付と時間を結合する処理
    state.value.dateTime = state.value.startDate + " ～ " + state.value.endDate;
  });
}
</script>
