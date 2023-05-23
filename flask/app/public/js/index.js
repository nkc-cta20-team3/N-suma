document.addEventListener('DOMContentLoaded', () => {

    LoadNavber()

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

// ナビゲーションバーを読み込む関数
function LoadNavber() {
    const navbar = document.getElementById('navber')
    
    const navbarContent = `
<nav class="navbar" role="navigation" aria-label="main navigation">
    
    <div class="navbar-brand">
        
        <!-- ロゴ -->
        <a class="navbar-item" href="index.html">
            <img src="#" alt="logo" width="112" height="40">
        </a>
    
        <!-- ハンバーガーメニュー -->
        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
        </a>
    </div>
  
    <div id="navbarBasicExample" class="navbar-menu">
        
        <div class="navbar-start">
        
            <a class="navbar-item" href="document_form.html">
                各種書類提出
            </a>
        
            <a class="navbar-item" href="document_list.html">
                各種書類閲覧
            </a>
  
            <a class="navbar-item" href="#">
                個人データ
            </a>
            
            <a class="navbar-item" href="document_auth.html">
                書類認可
            </a>

        </div>
  
        <div class="navbar-end">
            <div class="navbar-item">
                <div class="buttons">
                    
                    <a class="button is-primary" href="login.html">
                        <strong>Log in</strong>
                    </a>
                    <a class="button is-light">
                        <strong>Sign out</strong>
                    </a>

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
    `

    navbar.innerHTML = navbarContent
}
