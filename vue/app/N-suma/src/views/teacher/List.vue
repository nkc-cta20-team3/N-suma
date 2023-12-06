<template>
  <v-container>
    <!-- 書類一覧 -->
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <template class="d-flex flex-row justify-end text-black">
          <!-- 学籍番号入力 -->
          <v-text-field
            v-model="number"
            label="学籍番号"
            persistent-hint
            placeholder="20230001"
            persistent-placeholder
            :counter="8"
            class="mr-2"
          ></v-text-field>

          <!-- 検索ボタン -->
          <v-btn @click="onSearch" icon>
            <!-- icon url : https://pictogrammers.com/library/mdi/icon/magnify/ -->
            <v-icon alt="search icon" :icon="mdiMagnify"></v-icon>
          </v-btn>
        </template>
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
import { mdiMagnify } from "@mdi/js";
import { onMounted, ref } from "vue";
import router from "@/router";
import{ APICall } from "@utils"

const items = ref([]);
const number = ref("");

function onSearch() {
  console.log("検索");
}

function onItemClick(item) {
  console.log(item);
  router.push({
    name: "teacherView",
    params: { id: item.id },
  });
}

onMounted(() => {
  console.log("mounted");
  // TODO: 書類の一覧を取得し、itemsに格納する
  /*
  const readdocument_url = "/api/teacher/rd";
  APICall("POST", readdocument_url, state);
  
  console.log(json);
  */
  items.value = [
    { id: 1, class: "CTA20", name: "山田太郎" },
    { id: 2, class: "CTB20", name: "鈴木花子" },
    { id: 3, class: "CTA21", name: "佐藤次郎" },
  ];
});
</script>
