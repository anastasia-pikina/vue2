<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios';

const message = ref('Привет, мир!')
const news = ref();
const isDownload = ref(false)
const updateMessage = () => {
  message.value = 'Сообщение обновлено!'
}
onMounted(() => {
  console.log('размонтирован!')
  axios.get('http://localhost:8081/news')
      .then(response => {
        news.value = response.data;
        isDownload.value = true;
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

const getNewUrl = (NewId) => {
  return ['/news', NewId].join('/') + '/';
}

</script>

<template>
  <div class="banner-2 space-y-10 pb-20" id="work" v-if="isDownload">
    <h3 class="heading3 my-5">Новости</h3>
    <div class="card" v-for="(newItem, code) in news" :key="code">
      <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2">
        <h4 class="project-title item">{{newItem.name}}</h4>
        <p class="font-work_sans pr-12" v-html="newItem.description"></p>
        <a class="text-sky-800 font-bold text-2xl tracking-wider" :href="getNewUrl(newItem.id)">
          Подробнее
        </a>
      </div>
      <div v-if="newItem.image" class="card-image bg-green-100">
        <img
            class="object-cover w-full h-72 md:h-96"
            :src="'http://localhost:8081' + newItem.image"
        />
      </div>
    </div>


  </div>
</template>

<style scoped>

</style>