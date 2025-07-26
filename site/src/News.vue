<script setup>
import {ref, onMounted, watch, computed} from 'vue'
import axios from 'axios';
import Pagination from "./pagination/Pagination.vue";
import { useRoute } from 'vue-router'

const news = ref();
const isDownload = ref(false)
const currentPage = ref(1);
const totalCount = ref(0);

onMounted(() => {
  const route = useRoute()
  if (route.params.page) {
    currentPage.value = route.params.page;
  }

  axios.get(`http://localhost:4000/news?page=${currentPage.value}&limit=2`)
      .then(response => {
        news.value = response.data.list;
        totalCount.value = response.data.count;
        isDownload.value = true;
      })
      .catch(error => {
        console.log(error);
      });
})

const getNewUrl = (NewId) => {
  return ['/news', NewId].join('/') + '/';
}

const changePage = (pageNumber) => {
  console.log(pageNumber)
  if (pageNumber <= 0 || pageNumber > pageCount.value) {
    return;
  }

  currentPage.value = pageNumber
  axios.get(`http://localhost:4000/news?page=${pageNumber}&limit=2`)
      .then(response => {
        news.value = response.data.list;
        isDownload.value = true;
      })
      .catch(error => {
        console.log(error);
      });
}

const pageCount = computed(() => {
  try {
    return Math.ceil(totalCount.value / 2);
  } catch (e) {
    console.error(e.message);
  }

  return 0;
});

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
            :src="'http://localhost:4000' + newItem.image"
        />
      </div>
    </div>
    <Pagination :currentPage="currentPage" :pageCount="pageCount" :visiblePagesCount="11"
                @changePage="changePage"
    ></Pagination>
  </div>
</template>

<style scoped>

</style>