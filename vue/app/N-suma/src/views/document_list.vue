<template>
  <v-app id="inspire">
      <!-- ヘッダーコンポーネントを読み込む -->
      <Header />

    <v-main>
      <!--  -->
      <v-card class="mx-auto" max-width="586">
        <v-list :items="item1" item-title="name" item-value="id"></v-list>
      </v-card>
      <v-card class="mx-auto" max-width="586">
        <v-list :items="item2" item-title="name" item-value="id"></v-list>
      </v-card>
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
    item1: [
      {
        name: "Item #1",
        id: 1,
      },
    ],
    item2: [
      {
        name: "Item #2",
        id: 2,
      },
    ],
  }),
};
</script>
