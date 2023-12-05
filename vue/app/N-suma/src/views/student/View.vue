<template>
  <v-container>
    <v-row justify="center">
      <!-- ひとつ前の書類を閲覧するボタン -->
      <v-col cols="1">
        <v-btn @click="onClick" height="100%" elevation="0" size="small">
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
        <RowCard title="申請時間" :text="dateTime" />
        <RowCard title="場所" :text="state.location" />
        <RowCard title="必要欠席時間" :text="absenceTime" />
        <RowCard title="学生コメント" :text="state.studentComment" />
        <RowCard title="教員コメント" :text="state.teacherComment" />
      </v-col>
      
      <!-- ひとつ後の書類を閲覧するボタン -->
      <v-col cols="1">
        <v-btn @click="onClick" height="100%" elevation="0" size="small">
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
import router from "@/router";
import RowCard from "@/components/StudentViewRowCard.vue";
import{APICall} from "@/utils";

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
  prevId: "",
  nextId: "",
});

var dateTime = "";
var absenceTime = "";

function onClick() {
  // TODO: ひとつ前の書類を閲覧するボタンを押したときの処理を記述する
  // TODO: ひとつ後の書類を閲覧するボタンを押したときの処理を記述する
  const urln = "/api/admin/nd";
  APICall("POST", urln, state);
  
  console.log(json);
  console.log("onClick");
}

onMounted(() => {
  console.log("mounted");
  state.value.id = router.currentRoute.value.params.id;
  console.log(state.value.id);
  // TODO: 受け取ったidをもとに、ドキュメントを取得するAPIを叩き、stateに格納する
  
  const urlr = "/api/admin/rd";
  APICall("POST", urlr, state);
  
  console.log(json);

  state.value = {
    date: "2021-04-01",
    division: "国家試験 / FE",
    startDate: "11-15 22:20",
    endDate: "11-15 22:20",
    location: "愛知県名古屋市",
    startAbsence: "1限目",
    endAbsence: "6限目",
    studentComment:
      "180文字 : ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
    teacherComment: "把握しました",
  };
  // TODO: 現在の書類の前後の書類を取得し、stateに格納する
  state.value.prevId = state.value.id - 1;
  state.value.nextId = state.value.id + 1;

  // TODO: 表示用に日付と時間を結合する処理を記述する
  dateTime = state.value.startDate + " ～ " + state.value.endDate;
  absenceTime = state.value.startAbsence + " ～ " + state.value.endAbsence;
});
</script>
