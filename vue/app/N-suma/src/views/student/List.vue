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
import { useStore } from "@/stores/user";
import { APICallonJWT } from "@/utils";

const store = useStore();
const items = ref([]);

function onItemClick(item) {
  // console.log(item);
  router.push({
    name: "studentView",
    params: { id: item.id },
  });
}

onMounted(() => {
  // 書類一覧を取得する処理
  APICallonJWT("student/list/read", {
    user_id: store.id,
  }).then((res) => {
    // console.log(res);
    res.document.forEach((docs) => {
      items.value.push({
        id: docs.document_id,
        division: docs.division_name + "/" + docs.division_detail,
        date: docs.request_at,
      });
    });
  });
});
</script>
