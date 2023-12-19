<template>
  <v-row class="d-flex flex-row text-black pt-2">
    <v-card width="35%">
      <v-card-title>{{ props.title }}</v-card-title>
    </v-card>
    <v-card width="65%" @click.stop="dialog = true">
      <v-card-title>{{ props.text }}</v-card-title>
    </v-card>
  </v-row>
  <v-dialog width="70%" v-model="dialog" scrollable>
    <v-card>
      <v-card-text>{{ props.text }}</v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { defineProps, ref, toRef } from "vue";
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
toRef(props, "text");

const dialog = ref(false);
</script>
