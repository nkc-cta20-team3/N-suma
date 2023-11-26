<template>
  <v-app id="inspire">
    <v-main>
      <!--  -->

      <v-row justify="center">
        <v-col cols="12" sm="10" md="8" lg="6">
          <v-card ref="form">
            <v-card-text>
              <v-text-field
                ref="name"
                v-model="name"
                :rules="[() => !!name || '必須']"
                :error-messages="errorMessages"
                label="氏名"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="student_number"
                v-model="student_number"
                :rules="[() => !!student_number || '必須']"
                :error-messages="errorMessages"
                label="学籍番号"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="indispensable"
                v-model="indispensable"
                :rules="[() => !!indispensable || '必須']"
                :error-messages="errorMessages"
                label="クラス略称"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="position"
                v-model="position"
                :rules="[() => !!position || '必須']"
                :error-messages="errorMessages"
                label="役職"
                placeholder=""
                required
              ></v-text-field>
            </v-card-text>
            <v-divider class="mt-12"></v-divider>
            <!-- ボタン処理 -->
            <v-card-actions>
              <!-- 更新ボタン -->
              <v-container fluid>
                <v-btn color="green" @click="update_showAlertDialog"
                  >更新</v-btn
                >
                <v-dialog v-model="update_dialog" max-width="400">
                  <v-card>
                    <v-card-title>ユーザーを更新しますか？</v-card-title>
                    <v-card-actions>
                      <v-btn @click="update_registerUser">更新</v-btn>
                      <v-btn @click="update_cancelRegistration"
                        >キャンセル</v-btn
                      >
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-container>

              <!-- 削除ボタン -->
              <v-container fluid>
                <v-btn color="red" @click="delete_showAlertDialog">削除</v-btn>
                <v-dialog v-model="delete_dialog" max-width="400">
                  <v-card>
                    <v-card-title>ユーザーを削除しますか？</v-card-title>
                    <v-card-actions>
                      <v-btn @click="delete_registerUser">削除</v-btn>
                      <v-btn @click="delete_cancelRegistration"
                        >キャンセル</v-btn
                      >
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-container>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    components: {
      Header,
    },
  }),

  data() {
    return {
      delete_dialog: false,
      update_dialog: false,
    };
  },
  methods: {
    // ユーザー削除の処理
    delete_showAlertDialog() {
      this.delete_dialog = true;
    },
    delete_registerUser() {
      console.log("ユーザーを削除しました");
      this.delete_dialog = false;
    },
    delete_cancelRegistration() {
      console.log("削除をキャンセルしました");
      this.delete_dialog = false;
    },

    // ユーザー更新の処理
    update_showAlertDialog() {
      this.update_dialog = true;
    },
    update_registerUser() {
      console.log("ユーザーを更新しました");
      this.update_dialog = false;
    },
    update_cancelRegistration() {
      console.log("更新をキャンセルしました");
      this.update_dialog = false;
    },
  },
};
</script>
