<template>
  <v-row class="d-flex flex-row text-black pt-2">
    <v-card width="100%" @click="onItemClick">
      <v-card-title>{{ props.title }} : {{ props.text }}</v-card-title>
    </v-card>
  </v-row>
</template>

<script setup>
import { defineProps, ref } from "vue";
import router from "@/router";
const dialog = ref(false);

const props = defineProps({
  id: {
    type: Number,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
});
function onItemClick() {
  // console.log(props.id);
  router.push({
    name: "studentReForm",
    params: { id: props.id },
  });
}
</script>
