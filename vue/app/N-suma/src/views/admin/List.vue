<template>
  <v-container>
    <!-- 検索バー -->
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6" class="pt-10">
        <v-form ref="mainForm" class="d-flex flex-row justify-end text-black">
          <!-- 役職選択 -->
          <v-select
            v-model="role"
            :items="roles"
            label="役職"
            persistent-hint
            persistent-placeholder
            :rules="requiredRules"
            class="mr-2"
          ></v-select>

          <!-- 学籍番号入力 -->
          <v-text-field
            v-if="role === '学生'"
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
        </v-form>
      </v-col>
    </v-row>

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
          <v-list-item v-if="items.length == 0">
            <v-list-item-title
              >該当するユーザーが見つかりませんでした</v-list-item-title
            >
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
import {
  requiredRules,
  numberRules,
  APICallonJWT,
  APICallonGET,
} from "@/utils";

const mainForm = ref(null);

// 役職一覧を格納する変数
var roles = ref([]);
var roleids = [];

const role = ref("");
const number = ref("");

const items = ref([]);

async function onSearch() {
  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("入力エラー");
    return;
  }

  // 学生でない場合は学籍番号を空にする
  if (role.value != "学生") {
    number.value = "";
  }

  // ユーザーを検索する処理
  APICallonJWT("admin/list/search", {
    post_id: roleids[roles.value.indexOf(role.value)],
    user_number: number.value,
  }).then((res) => {
    // console.log(res);
    items.value = [];
    res.document.forEach((user) => {
      items.value.push({
        id: user.user_id,
        class: user.class_abbr == "" ? "None" : user.class_abbr,
        name: user.user_name,
      });
    });
  });
}

function onItemClick(item) {
  // console.log(item);
  router.push({
    name: "adminEdit",
    params: { id: item.id },
  });
}

onMounted(() => {
  init();
});

function init() {
  // ロール一覧を取得する
  APICallonGET("utils/read/post").then((res) => {
    // console.log(res);
    res.document.forEach((post) => {
      roles.value.push(post.post_name);
      roleids.push(post.post_id);
    });
  });

  // ユーザー一覧を取得する処理
  APICallonJWT("admin/list/read", {}).then((res) => {
    // console.log(res);
    res.document.forEach((user) => {
      items.value.push({
        id: user.user_id,
        class: user.class_abbr == "" ? "None" : user.class_abbr,
        name: user.user_name,
      });
    });
  });
}
</script>
