<template>
  <v-row class="d-flex flex-row text-black pt-2">
    <v-card width="100%">
      <v-card-title>{{ props.title }} : {{ props.text }}</v-card-title>
    </v-card>
  </v-row>
</template>

<script setup>
import { defineProps } from "vue";

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
});
</script>
