<template>
  <v-app id="inspire">
    <!-- ヘッダーコンポーネントを読み込む -->
    <Header />
    <!-- <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title><Nav>Nスマ</Nav></v-toolbar-title>
    </v-app-bar> -->

    <v-main>
      <!--  -->

      <v-row justify="center">
        <v-col cols="12" sm="10" md="8" lg="6">
          <v-card ref="form">
            <v-card-text>
              <v-autocomplete
                ref="section"
                v-model="section"
                :rules="[() => !!section || '必須項目']"
                :items="sections"
                label="区分"
                placeholder="選択"
                required
              ></v-autocomplete>
              <v-text-field
                ref="StartTime"
                v-model="StartTime"
                :rules="[() => !!StartTime || '必須']"
                :error-messages="errorMessages"
                label="公欠開始時間"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="EndTime"
                v-model="EndTime"
                :rules="[() => !!EndTime || '必須']"
                :error-messages="errorMessages"
                label="公欠終了時間"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="location"
                v-model="location"
                :rules="[() => location || '必須']"
                :error-messages="errorMessages"
                label="実施場所"
                placeholder=""
                required
              ></v-text-field>
              <v-text-field
                ref="comment"
                v-model="comment"
                :rules="[() => !!comment || '必須']"
                :error-messages="errorMessages"
                label="コメント"
                placeholder=""
                required
              ></v-text-field>
            </v-card-text>
            <v-divider class="mt-12"></v-divider>
            <v-card-actions>
              <v-btn variant="text"> Cancel </v-btn>
              <v-spacer></v-spacer>
              <v-slide-x-reverse-transition>
                <v-tooltip v-if="formHasErrors" location="left">
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn
                      icon
                      class="my-0"
                      v-bind="attrs"
                      @click="resetForm"
                      v-on="on"
                    >
                      <v-icon>mdi-refresh</v-icon>
                    </v-btn>
                  </template>
                  <span>Refresh form</span>
                </v-tooltip>
              </v-slide-x-reverse-transition>
              <v-btn color="primary" variant="text" @click="submit">
                提出
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-main>
  </v-app>
</template>

<script>
//ヘッダー読み込み用
import Header from '../components/Navigation.vue';

export default {
  data: () => ({
    components:{
      Header,
    },
    sections: ["就活", "資格試験", "弔事"],
    errorMessages: "",
    name: null,
    address: null,
    city: null,
    state: null,
    zip: null,
    country: null,
    formHasErrors: false,
  }),

  computed: {
    form() {
      return {
        name: this.name,
        address: this.address,
        city: this.city,
        state: this.state,
        zip: this.zip,
        country: this.country,
      };
    },
  },

  watch: {
    name() {
      this.errorMessages = "";
    },
  },

  methods: {
    addressCheck() {
      this.errorMessages =
        this.address && !this.name ? `Hey! I'm required` : "";

      return true;
    },
    resetForm() {
      this.errorMessages = [];
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        this.$refs[f].reset();
      });
    },
    submit() {
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        if (!this.form[f]) this.formHasErrors = true;

        this.$refs[f].validate(true);
      });
    },
  },
};
</script>
