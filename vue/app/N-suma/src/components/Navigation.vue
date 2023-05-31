<template>

    <nav class="navbar" role="navigation" aria-label="main navigation">    
        <div class="navbar-brand">
            
            <!-- ロゴ -->
            <router-link class="navbar-item" to="/">
                <img src="#" alt="logo" width="112" height="40">
            
                <img src="/vite.svg" width="112" height="28">
            </router-link>

            <!-- ハンバーガーメニュー -->
            <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
            </a>
        </div>
    
        <div id="navbarBasicExample" class="navbar-menu">
            
            <div class="navbar-start">
            
                <router-link class="navbar-item" to="/document_form">
                    各種書類提出
                </router-link>
            
                <router-link class="navbar-item" to="/document_list">
                    各種書類閲覧
                </router-link>
    
                <router-link class="navbar-item" to="#">
                    個人データ
                </router-link>
                
                <router-link class="navbar-item" to="/document_auth">
                    書類認可
                </router-link>

            </div>
    
            <div class="navbar-end">
                <div class="navbar-item">
                    <div class="buttons">
                        
                        <router-link class="button is-primary" to="/login">
                            <strong>Log in</strong>
                        </router-link>
                        <router-link class="button is-light">
                            <strong>Sign out</strong>
                        </router-link>

                        <span class="material-symbols-outlined">
                            notifications
                        </span>
                        <span class="material-symbols-outlined">
                            account_circle
                        </span>
                    </div>
                </div>
            </div>

        </div>
    </nav>

</template>

<script setup>
    import router from '../router';

    document.addEventListener('DOMContentLoaded', () => {
        // Get all "navbar-burger" elements
        const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0)
        
        $navbarBurgers.forEach( el => {
            el.addEventListener('click', () => {
                const target = el.dataset.target
                const $target = document.getElementById(target)
                el.classList.toggle('is-active')
                $target.classList.toggle('is-active')
            })
        })
    })

</script>

