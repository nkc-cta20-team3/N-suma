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
            
                <router-link class="navbar-item" to="/" v-if="!isLoggedIn">
                    Home
                </router-link>

                <router-link class="navbar-item" to="/document_form" v-if="isLoggedIn">
                    各種書類提出
                </router-link>
            
                <router-link class="navbar-item" to="/document_list" v-if="isLoggedIn">
                    各種書類閲覧
                </router-link>
    
                <router-link class="navbar-item" to="#" v-if="isLoggedIn">
                    個人データ
                </router-link>
                
                <router-link class="navbar-item" to="/document_auth" v-if="isLoggedIn">
                    書類認可
                </router-link>

            </div>
    
            <div class="navbar-end">
                <div class="navbar-item">
                    <div class="buttons">
                        
                        <a @click="signInWithGoogle" v-if="!isLoggedIn" class="button is-primary">    
                            <strong>Log in</strong>
                        </a>
                        <a @click="handleSignOut" v-if="isLoggedIn" class="button is-light">
                            <strong>Sign Out</strong>
                        </a>

                        <span class="material-symbols-outlined" v-if="isLoggedIn">
                            notifications
                        </span>
                        <span class="material-symbols-outlined" v-if="isLoggedIn">
                            account_circle
                        </span>
                    </div>
                </div>
            </div>

        </div>
    </nav>

</template>

<script setup>
    import { onMounted, ref } from 'vue'
    import { getAuth, onAuthStateChanged, signOut , GoogleAuthProvider, signInWithPopup } from 'firebase/auth'
    import router from '../router';

    const isLoggedIn = ref(false)

    let auth
    onMounted(() => {
        auth = getAuth()
        onAuthStateChanged(auth, (user) => {
            isLoggedIn.value = !!user
        })
    })

    const handleSignOut = () => {
        signOut(auth).then(() => {
            router.push('/')
            isLoggedIn.value = false
        }).catch((error) => {
            console.log(error)
        })
    }

    const signInWithGoogle = () => {

        const provider = new GoogleAuthProvider()
        signInWithPopup(getAuth(), provider)
            .then((result) => {
                //console.log(result.user)
                router.push('/document_form')
            }).catch((error) => {
                console.log(error)
            })

    }

    //GetPotisionAPIに引数userIdを渡し返り値potisionIDを受け取る
   
const userId = 0;
      fetch(new URL(`fakeendpoint`, import.meta.env.VITE_API_URL))
      .then((response) => {
          if (!response.ok) {
             throw new Error(`${response.status} ${response.statusText}`)
          }
          return response.json()
      })
      .then((data) => {
          posts.value = data.posts; // 取得したデータのpostsを格納
      })
      .catch((error) => {
          console.log(error)
      })


    //ハンバーガーメニューの実装
    document.addEventListener('DOMContentLoaded', () => {

        // トグルボタンを取得
        const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

        // トグルボタンが存在する場合
        if ($navbarBurgers.length > 0) {

            // トグルボタンにクリックイベントを設定
            $navbarBurgers.forEach( el => {
                el.addEventListener('click', () => {

                    // data-targetの属性値からナビゲーションメニューを取得
                    const target = el.dataset.target;
                    const $target = document.getElementById(target);

                    // トグルボタンとナビゲーションメニューにis-activeクラスを設定
                    el.classList.toggle('is-active');
                    $target.classList.toggle('is-active');

                });
            });
        }
    });

</script>

