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

        <div class="is-grouped">
            
            <form method="post" action="" class="form-horizontal">
                <div class="form-group">

                    <div class="col-sm-4">
                        <input type="text" id="name" name="id" class="form-control" required="required" placeholder="教員用コメント">
                    </div>
                </div>
                <input type="submit" class="button" value="受理">
                <input type="submit" class="button" value="却下">
            </form>

        </div>

    </div>
</template>