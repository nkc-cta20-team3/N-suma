<template>
    <Navigation />
    <router-view />
</template>

<script setup>
    import Navigation from './components/Navigation.vue'

</script>

<style>
    @import 'bulma/css/bulma.css';
    th {
    width: 150px;
}

.container{
    display: flex;
    justify-content: center;
}

</style>