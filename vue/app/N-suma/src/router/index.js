import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {  path: '/', component: () => import('../views/index.vue') },
        {  path: '/login', component: () => import('../views/login.vue') },
        {  path: '/document_list', component: () => import('../views/document_list.vue') },
        {  path: '/document_form', component: () => import('../views/document_form.vue') },
        {  path: '/document_auth', component: () => import('../views/document_auth.vue') },
        {  path: '/document_accept', component: () => import('../views/document_accept.vue') },
        
    ],
})

export default router