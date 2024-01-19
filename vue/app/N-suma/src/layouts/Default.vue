<template>
  <v-app>
    <Navigation />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import Navigation from "@/components/Navigation.vue";
</script>
