<template>
  <v-container>
    <!-- 書類一覧 -->
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <v-form ref="mainForm" class="d-flex flex-row justify-end text-black">
          <!-- 学籍番号入力 -->
          <v-text-field
            v-model="number"
            label="学籍番号"
            persistent-hint
            placeholder="20201001"
            persistent-placeholder
            :rules="numberRules"
            :counter="8"
            class="mr-2"
          ></v-text-field>

          <!-- 検索ボタン -->
          <v-btn @click="onSearch" icon>
            <!-- icon url : https://pictogrammers.com/library/mdi/icon/magnify/ -->
            <v-icon alt="search icon" :icon="mdiMagnify"></v-icon>
          </v-btn>

          <v-btn @click="getAllDocs" icon>
            <!-- icon url : https://pictogrammers.com/library/mdi/icon/reload/ -->
            <v-icon alt="search icon" :icon="mdiReload"></v-icon>
          </v-btn>
        </v-form>

        <v-list>
          <v-list-item
            v-for="item in items"
            :key="item.id"
            @click="onItemClick(item)"
            :elevation="2"
            height="60"
            class="mb-2"
          >
            <v-list-item-title>
              {{ item.class }} : {{ item.name }}
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { mdiMagnify, mdiReload } from "@mdi/js";
import router from "@/router";
import { numberRules, APICallonJWT } from "@/utils";

const mainForm = ref(null);

const number = ref("");
const items = ref([]);

async function onSearch() {
  // TODO: 学籍番号を元に検索し、itemsに格納する
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // ユーザーを検索する処理
  APICallonJWT("teacher/viewlist/search", {
    user_number: number.value,
  }).then((res) => {
    // console.log(res);
    items.value = [];
    res.document.forEach((docs) => {
      items.value.push({
        id: docs.document_id,
        class: docs.division_name + "/" + docs.division_detail,
        name: docs.user_name,
      });
    });
  });
}

function onItemClick(item) {
  // console.log(item);
  router.push({
    name: "teacherView",
    params: { id: item.id },
  });
}

onMounted(() => {
  getAllDocs();
});

function getAllDocs() {
  // 書類一覧を取得する処理
  APICallonJWT("teacher/viewlist/read", {}).then((res) => {
    // console.log(res);
    items.value = [];
    res.document.forEach((docs) => {
      items.value.push({
        id: docs.document_id,
        class: docs.division_name + "/" + docs.division_detail,
        name: docs.user_name,
      });
    });
  });
}
</script>
