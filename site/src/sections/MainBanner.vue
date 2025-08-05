<script setup>
import { ref, onMounted } from 'vue'
import axios from "axios";
const contentBanner = ref();
const isDownload = ref(false)

onMounted(() => {
  axios.get('http://localhost:4000/blocks/main')
      .then(response => {
        contentBanner.value = response.data.find(x=>x!==undefined);
        isDownload.value = true;
      })
      .catch(error => {
        console.log(error);
      });
})
</script>

<template>
  <div class="banner-2 md:flex items-center md:justify-evenly" id="about" v-if="isDownload">
    <img
        class="md:h-[500px] h-[400px]"
        :src="'http://localhost:4000' + contentBanner.image"
        alt="user image"
    />
    <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2">
      <h4 class="project-title item">{{contentBanner.title}}</h4>
      {{contentBanner.content}}
    </div>
  </div>
</template>

<style scoped>

</style>