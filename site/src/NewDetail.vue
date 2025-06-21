<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios';

const message = ref('Привет, мир!')
const news = ref();
const updateMessage = () => {
  message.value = 'Сообщение обновлено!'
}
onMounted(() => {
  console.log('размонтирован!')
  axios.get('http://shop2.local:8081/users')
      .then(response => {
        news.value = response.data;
        console.log(response.data)
      })
      .catch(error => {
        console.log(error);
      });
})

const getImage = (fileName, subDir) => {
  if (!fileName.Valid) {
    return '';
  }

  return ['http://shop2.local/upload/', subDir.String, fileName.String].join('/');
}


// Нет необходимости в return, все переменные и функции автоматически доступны в шаблоне
</script>

<template>
  New Detail
</template>

<style scoped>

</style>