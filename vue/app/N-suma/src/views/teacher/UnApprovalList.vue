<template>
  <v-container>
    <!-- 書類一覧 -->
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
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
import router from "@/router";
import{APICall} from "@/utils";

const items = ref([]);

function onItemClick(item) {
  console.log(item);
  router.push({
    name: "auth",
    params: { id: item.id },
  });
}

onMounted(() => {
  console.log("mounted");
  // TODO: 書類の一覧を取得し、itemsに格納する
  /*
  const readauthlist_url = "/api/teacher/ral";
  APICall("POST", readauthlist_url, state);
  
  console.log(json);
*/

  items.value = [
    { id: 1, class: "CTA20", name: "山田太郎" },
    { id: 2, class: "CTB20", name: "鈴木花子" },
    { id: 3, class: "CTA21", name: "佐藤次郎" },
  ];
});
</script>
