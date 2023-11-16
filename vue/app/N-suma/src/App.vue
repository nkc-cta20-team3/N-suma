<template>
  <v-app>
    <Navigation />
    <router-view />
  </v-app>
</template>

<script setup>
import Navigation from "./components/Navigation.vue";
import { app } from "./firebase/index.js";
</script>

<style>
@import "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200";

th {
  width: 150px;
}

.container {
  display: flex;
  justify-content: center;
}
</style>
