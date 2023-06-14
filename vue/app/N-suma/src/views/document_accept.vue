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

<script setup>
    
    import { onMounted,ref } from 'vue'

    const APIURL = new URL("ad/1", import.meta.env.VITE_API_URL)

    onMounted(() => {
        GetAPIDate() 
    })

    const GetAPIDate = () => {
        fetch(APIURL, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'X-Requested-With': 'XMLHttpRequest',
            },
        })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`${response.status} ${response.statusText}`)
            } else {
                response.json().then((data) => {
                    console.log(data)
                })
            }
        })
        .then((data) => {
            console.log(data)
        })
        .catch((error) => {
            console.log(error)
        })
    }

</script>