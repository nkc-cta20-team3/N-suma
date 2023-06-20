<template>
    <div class="title has-text-centered">
        各種書類認可
    </div>

    <div class="columns">
        
        <div class="column has-text-right">
            <div>未認可リスト</div>
        </div>
        
        <div class="column"></div>

        <div class="column">
            <div>ソート</div>
            <select class="select">
                <option>クラス略称</option>
                <option>名前</option>
                <option>分類</option>
            </select>
        </div>

    </div>

    <div class="container">
        <table class="table">
            <thead>
                <tr>
                    <th>クラス略称</th>
                    <th>氏名</th>
                    <th>分類</th>
                    <th>    </th>
                </tr>
            </thead>
            <thead>
                <tr v-for="user in users" :key="user.id">
                    <th>{{ user.class }}</th>
                    <th>{{ user.name }}</th>
                    <th>{{ user.type }}</th>
                    <th><button @click ="showUserDetail(user)">詳細</button></th>
                </tr>
            </thead>
        </table>
    </div>
</template>

<script setup>

    import store from '../store'
    import { useRouter } from 'vue-router'
    import { onMounted,ref } from 'vue'

    const router = useRouter();

    const users = ref(
        [
            {
                id: 1,
                class: 'CS2',
                name: '石井 大介',
                type: '就活',
            },
                {
                id: 2,
                class: 'CT4B',
                name: '鈴木 孝明',
                type: '就活',
            },
                {
                id: 3,
                class: 'CT2A',
                name: '遠藤 雄一',
                type: '資格',
            },
        ]
    )

    const showUserDetail = (user) => {
        router.push({ name: 'document_accept', params: {userId: user.class}});
        store.commit('setClassId',)
    }

    onMounted(() => {
        if(users.value.length < 1){
            alert('未認可の書類はありません。');
        }
        else{
            
        }    
    })


</script>