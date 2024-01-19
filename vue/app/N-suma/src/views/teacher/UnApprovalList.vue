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
import { APICallonJWT } from "@/utils";

const items = ref([]);

function onItemClick(item) {
  // console.log(item);
  router.push({
    name: "auth",
    params: { id: item.id },
  });
}

onMounted(() => {
  // 書類一覧を取得する処理
  APICallonJWT("teacher/unauthllist/read", {}).then((res) => {
    // console.log(res);
    res.document.forEach((docs) => {
      items.value.push({
        id: docs.document_id,
        class: docs.division_name + "/" + docs.division_detail,
        name: docs.user_name,
      });
    });
  });
});
</script>
