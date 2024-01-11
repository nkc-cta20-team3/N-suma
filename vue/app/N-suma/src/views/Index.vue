<template>
  <!-- キービジュアルを置く -->
  <v-container>
    <v-row justify="center">
      <v-col cols="12" sm="10" md="8" lg="6">
        <p
          class="display-1 text-center"
          v-if="store.getIsLogin && !store.getRole"
        >
          あなたのアカウントは利用できません。<br />
          役職が未割当か、凍結されている可能性があります。<br />
          管理者に問い合わせてください。
        </p>
        <v-img src="/Poster.gif" height="100%" width="100%"></v-img>
      </v-col>
    </v-row>
  </v-container>
</template>
<script setup>
import { useStore } from "@/stores/user";
const store = useStore();
</script>
