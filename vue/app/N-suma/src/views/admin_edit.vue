<template>
  <v-app id="inspire">
    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title><Nav>Nスマ</Nav></v-toolbar-title>
    </v-app-bar>

    <v-main>
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
    </v-main>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    sections: ["担任", "学生"],
    errorMessages: "",
    name: null,
    address: null,
    city: null,
    state: null,
    zip: null,
    country: null,
    formHasErrors: false,
  })
};
</script>