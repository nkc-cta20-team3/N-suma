<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <!-- 書類一覧 -->
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
              {{ item.date }} : {{ item.division }}
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";
import router from "@/router";
import {APICall} from "@/utils";

const items = ref([]);

function onItemClick(item) {
  console.log(item);
  router.push({
    name: "studentView",
    params: { id: item.id },
  });
}

onMounted(() => {
  console.log("mounted");
  // TODO: ユーザーごとの書類の一覧を取得し、itemsに格納する
  /*
  const url = "/api/admin/ral";
  APICall("POST", url, state);
  
  console.log(json);
*/
  items.value = [
    { id: 1, division: "国家試験 / FE", date: "2021/04/01" },
    { id: 2, division: "国家試験 / SG", date: "2021/04/01" },
    { id: 3, division: "国家試験 / AP", date: "2021/04/01" },
  ];
});
</script>
