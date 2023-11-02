<template>
<div v-if = "userId === 'admin'">
    <h1>管理者編集画面だお</h1>
</div>
<div iv-if="userId != 'admin'">
<h1>不法侵入を検知　ただちにもとのページに戻りなさい</h1>
</div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { getAuth, onAuthStateChanged, signOut , GoogleAuthProvider, signInWithRedirect } from 'firebase/auth'
import router from '../router';
//import store from './store';


let auth
onMounted(() => {
    auth = getAuth()
    onAuthStateChanged(auth, (user) => {
        isLoggedIn.value = !!user

        if(user){
            store.dispatch('updateUser', {
                    userId: userId,
                    classId: ClassId,
                    DocId: DocId
            })
        }
    })
})



</script>