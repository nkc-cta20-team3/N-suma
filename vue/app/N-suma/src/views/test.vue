<template>
    <div>
      <button @click="fetchData">APIを呼び出す</button>
      <div>{{ responseData }}</div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        responseData: ''
      };
    },
    methods: {
      fetchData() {
        fetch('http://localhost:8080/api/ual/1') // APIのエンドポイントを指定
          .then(response => {
            if (!response.ok) {
              throw new Error('APIリクエストが失敗しました。');
            }
            return response.json();
          })
          .then(data => {
            this.responseData = data; // レスポンスデータを表示するためにresponseDataに代入
          })
          .catch(error => {
            console.error(error);
          });
      }
    }
  };
  </script>
  
