<template>
  <v-container>
    <v-row justify="center">
      <!-- ひとつ前の書類を閲覧するボタン -->
      <v-col cols="1">
        <v-btn
          v-if="prevId != Number(-1)"
          @click="onPrevClick"
          height="100%"
          elevation="0"
          size="small"
        >
          <!-- icon url : https://pictogrammers.com/library/mdi/icon/menu-left/ -->
          <v-icon :icon="mdiMenuLeft" size="x-large"></v-icon>
          <v-icon :icon="mdiMenuLeft" size="x-large"></v-icon>
          <v-icon :icon="mdiMenuLeft" size="x-large"></v-icon>
        </v-btn>
      </v-col>

      <!-- 書類詳細 -->
      <v-col cols="10">
        <RowCard title="申請日" :text="state.date" />
        <RowCard title="公欠区分" :text="state.division" />
        <RowCard title="申請時間" :text="state.dateTime" />
        <RowCard title="必要欠席時間" :text="state.absenceTime" />
        <RowCard title="場所" :text="state.location" />
        <RowCard title="学生コメント" :text="state.studentComment" />
        <RowCard title="教員コメント" :text="state.teacherComment" />
      </v-col>

      <!-- ひとつ後の書類を閲覧するボタン -->
      <v-col cols="1">
        <v-btn
          v-if="nextId != Number(-1)"
          @click="onNextClick"
          height="100%"
          elevation="0"
          size="small"
        >
          <!-- icon url : https://pictogrammers.com/library/mdi/icon/menu-right/ -->
          <v-icon :icon="mdiMenuRight" size="x-large"></v-icon>
          <v-icon :icon="mdiMenuRight" size="x-large"></v-icon>
          <v-icon :icon="mdiMenuRight" size="x-large"></v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { mdiMenuLeft, mdiMenuRight } from "@mdi/js";
import { onBeforeRouteUpdate } from "vue-router";
import router from "@/router";
import RowCard from "@/components/StudentViewRowCard.vue";
import { APICallonJWT } from "@/utils";

const state = ref({
  id: "",
  date: "",
  division: "",
  startDate: "",
  endDate: "",
  location: "",
  startAbsence: "",
  endAbsence: "",
  studentComment: "",
  teacherComment: "",
  dateTime: "",
  absenceTime: "",
});
const prevId = ref(-1);
const nextId = ref(-1);

function onPrevClick() {
  router.push({
    name: "teacherView",
    params: { id: prevId.value },
  });
}

function onNextClick() {
  router.push({
    name: "teacherView",
    params: { id: nextId.value },
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
  // console.log(state.value.id);

  // 書類の詳細情報を取得する
  APICallonJWT("teacher/view/read", {
    document_id: Number(state.value.id),
  }).then((res) => {
    // console.log(res);
    state.value = {
      date: res.document.request_at,
      division: res.document.division_name + "/" + res.document.division_detail,
      startDate: res.document.start_time,
      endDate: res.document.end_time,
      startAbsence: res.document.start_flame,
      endAbsence: res.document.end_flame,
      location: res.document.location,
      studentComment: res.document.student_comment,
      teacherComment: res.document.teacher_comment,
    };

    // 表示用に日付と時間を結合する処理
    state.value.dateTime = state.value.startDate + " ～ " + state.value.endDate;
    state.value.absenceTime =
      state.value.startAbsence + "限目 ～ " + state.value.endAbsence + "限目";
  });

  // 現在の書類の前後の書類IDを取得し、stateに格納する
  APICallonJWT("teacher/view/next", {
    document_id: Number(state.value.id),
  }).then((res) => {
    // console.log(res);
    nextId.value = res.document.document_id_next;
    prevId.value = res.document.document_id_prev;
  });
}
</script>
