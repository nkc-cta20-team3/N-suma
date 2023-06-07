<template>
    <div class="title has-text-centered">
        書類受理
    </div>

    <a class="has-text-left" style="font-size: 1.5em" href="document_auth.html">
        ◀戻る
    </a>

    <div class="has-text-centered">
        <iframe src="https://www.denpa.ac.jp/school_guide/pdf/02_CT3A.pdf" width="630" height="800" frameborder="0" marginheight="0" marginwidth="0">読み込んでいます…</iframe>
    </div>

    <div class="container">

        <form class="is-grouped">

            <div class="is-flex mb-5">
                <input type="text" disabled class="input disable" style="font-size: 1.25em" placeholder="学生コメント" >
            </div>

            <div class="is-flex mb-5">
                <input type="text" required class="input" style="font-size: 1.25em" placeholder="教員用コメント" >
            </div>
            
            <div class="is-flex mb-5">
                <input type="submit" class="ml-4 button is-medium" value="受理">
                <input type="submit" class="ml-4 button is-medium" value="却下">
            </div>

        </form>
    </div>
</template>