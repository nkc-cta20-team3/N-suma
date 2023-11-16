<template>
  <v-app id="inspire">
    <!-- ヘッダーコンポーネントを読み込む -->
    <Header />

    <v-main>
      <!-- 検索バー -->
      <v-row justify="center">
        <v-col cols="13" sm="10" md="8" lg="6">
          <v-card ref="form">
            <v-container>
              <v-row justify="center">
                <!-- テキストボックス -->
                <v-col cols="8">
                  <v-text-field
                    label="学籍番号"
                    v-model="studentId"
                    ref="student_number"
                  ></v-text-field>
                </v-col>

                <!-- 検索ボタン -->
                <v-col cols="1" class="d-flex justify-end align-center">
                  <v-btn @click="buttonClick" icon>
                    <v-icon>mdi-arrow-down-bold-circle</v-icon>
                  </v-btn>
                </v-col>

                <!-- 役職メニュー -->
                <v-col cols="2">
                  <v-autocomplete
                    ref="section"
                    v-model="section"
                    :items="sections"
                    label="区分"
                    placeholder="選択"
                    required
                    persistent-hint
                  ></v-autocomplete>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
      </v-row>

      <!-- 表示一覧 -->
      <v-row justify="center">
        <v-col>
          <router-link to="/admin_update">
            <v-card class="mx-auto" max-width="700">
              <v-list :items="item1" item-title="name" item-value="id"></v-list>
            </v-card>
          </router-link>
          <v-card class="mx-auto" max-width="700">
            <v-list :items="item2" item-title="name" item-value="id"></v-list>
          </v-card>
        </v-col>
      </v-row>
    </v-main>
  </v-app>
</template>

<script>
//ヘッダー読み込み用
import Header from "../components/Navigation.vue";
import router from "../router";

export default {
  data: () => ({
    components: {
      Header,
    },
    studentId: null,
    section: null,
    sections: ["担任", "学生"],
    item1: [
      {
        name: "学生1",
        id: 1,
      },
    ],
    item2: [
      {
        name: "学生2",
        id: 2,
      },
    ],
  }),
};
</script>
