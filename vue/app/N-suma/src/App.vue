<template>
    <Navigation />
    <router-view />
</template>

<script setup>
    import Navigation from './components/Navigation.vue'

</script>

<style>
    @import 'bulma/css/bulma.css';
    @import 'https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200';

    th {
        width: 150px;
    }

    .container{
        display: flex;
        justify-content: center;
    }

</style>