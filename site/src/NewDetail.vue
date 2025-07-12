<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router'

const message = ref('Привет, мир!')
const isDownload = ref(false)
const newDetail = ref();
const updateMessage = () => {
  message.value = 'Сообщение обновлено!'
}
onMounted(() => {
  const route = useRoute()
  const id = route.params.id
  axios.get('http://localhost:8081/news/' + id)
      .then(response => {
        newDetail.value = response.data;
        isDownload.value = true;
        console.log(response.data)
      })
      .catch(error => {
        console.log(error);
      });
})
</script>
<template>
  <div class="banner-2 space-y-10 pb-20" id="work" v-if="isDownload">
    <h3 class="heading3 my-5">Новости</h3>
    <div class="card">
      <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2">
        <h4 class="project-title item">{{newDetail.name}}</h4>
        <p class="font-work_sans pr-12" v-html="newDetail.content"></p>
      </div>
      <div v-if="newDetail.image" class="card-image bg-green-100">
        <img
            class="object-cover w-full h-72 md:h-96"
            :src="'http://localhost:8081' + newDetail.image"
        />
      </div>
    </div>
  </div>
</template>
<style scoped>
</style>