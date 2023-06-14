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
            <div class="is-flex mb-5">
                <input type="text" disabled class="input disable" style="font-size: 1.25em" placeholder="学生コメント" >
            </div>
            <div class="is-flex mb-5">
                <input type="text" v-model="inputValue" class="input" style="font-size: 1.25em" placeholder="教員用コメント">
            </div>
            <div class="is-flex mb-5">
                <div class="ml-4 button is-medium" @click="AcceptanceAlert">受理</div>
                <div class="ml-4 button is-medium" @click="RejectedAlert">却下</div>
            </div>
        </div>
    </div>

</template>
  
<script>
export default {
    data() {
        return {
            inputValue:""
        };
    },
    methods: {
        AcceptanceAlert() {
            if (this.inputValue){
                alert("受理されました");
            } else {
                alert("コメントを入力してください。");
            }      
        },
        RejectedAlert() {
            if (this.inputValue){
                alert("却下されました");
            } else {
                alert("コメントを入力してください。");
            }
        }
    }
};
</script>








