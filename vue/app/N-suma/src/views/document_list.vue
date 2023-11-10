<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title><Nav>Nスマ</Nav></v-toolbar-title>
    </v-app-bar>

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
export default {
  data: () => ({
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
