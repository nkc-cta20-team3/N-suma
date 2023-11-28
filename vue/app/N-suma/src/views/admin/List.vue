<template>
  <!-- 表示一覧 -->
  <v-row justify="center">
    <v-col>
      <router-link to="/admin_update">
        <v-card class="mx-auto" max-width="700">
          <v-list :items="item1" item-title="name" item-value="id"></v-list>
        </v-card>
      </router-link>
      <v-card>
        <v-list :items="item2" item-title="name" item-value="id"></v-list>
      </v-card>
    </v-col>
  </v-row>

  <v-container>
    <v-row justify="center">
      <!-- 検索バー -->
      <v-col cols="12" sm="10" md="8" lg="6">
        <template class="d-flex flex-row justify-end text-black">
          <!-- 役職選択 -->
          <v-select
            v-model="role"
            :items="roles"
            label="役職"
            persistent-hint
            class="pr-2"
          ></v-select>

          <!-- 学籍番号入力 -->
          <v-text-field
            v-model="number"
            label="学籍番号"
            persistent-hint
            placeholder="20201001"
            persistent-placeholder
            class="pr-2"
          ></v-text-field>
          <!-- 検索ボタン -->
          <v-btn @click="buttonClick" icon>
            <!-- font url : https://pictogrammers.com/library/mdi/icon/magnify/ -->
            <v-icon>mdi mdi-magnify</v-icon>
          </v-btn>
        </template>
      </v-col>

      <v-col cols="12" sm="10" md="8" lg="6">
        <v-list lines="one" v-for="item in items" :key="item.id">
          <v-hover v-slot:default="{ hover }">
            <v-expand-transition>
              <v-overlay absolute :opacity="0.2" :value="hover">
                <v-btn color="white" class="black--text"> Show Message </v-btn>
              </v-overlay>
            </v-expand-transition>
            <v-list-item class="border">
              <v-list-item-title>{{ item.name }}</v-list-item-title>
            </v-list-item>
          </v-hover>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { onMounted, ref } from "vue";

const roles = ["担任", "学生"];

const mainForm = ref(null);
const number = ref(null);
const role = ref(null);

const studentId = ref(null);
const section = ref(null);
const sections = ref(["担任", "学生"]);
const items = ref([
  {
    name: "学生1",
    id: 1,
  },
  {
    name: "学生1",
    id: 1,
  },
  {
    name: "学生1",
    id: 1,
  },
]);

async function buttonClick() {
  console.log("button clicked");

  const validResult = await mainForm.value.validate();
  if (!validResult.valid) {
    console.log("ユーザーを登録できませんでした");
    return;
  }
}

onMounted(() => {
  console.log("mounted");
});
</script>
