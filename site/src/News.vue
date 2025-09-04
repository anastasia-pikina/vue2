<script setup>
import {ref, onMounted, computed} from 'vue';
import Pagination from './pagination/Pagination.vue';
import { useRoute } from 'vue-router';
import Loader from './sections/Loader.vue';
import {useMainStore} from './stores/mainStore';
const news = ref([]);
const isDownloadProcess = ref(false);
const currentPage = ref(1);
const totalCount = ref(0);
const store = useMainStore();
const limit = 5;

onMounted( () => {
  try {
    const route = useRoute()
    if (route.params.page) {
      currentPage.value = route.params.page;
    }
    getNews();
  } catch (error) {
    console.error(error);
  }
})

const getNewUrl = (NewId) => {
  return ['/news', NewId].join('/') + '/';
}

const changePage = (pageNumber) => {
  try {
    if (pageNumber <= 0 || pageNumber > pageCount.value) {
      return;
    }

    currentPage.value = pageNumber;
    getNews();
  } catch (error) {
    console.error(error);
  }
}

const getNews = async () => {
  isDownloadProcess.value = true;
  const newsData = await store.getData(`/news?page=${currentPage.value}&limit=${limit}`);

  if (newsData) {
    news.value = newsData.list ?? [];
    totalCount.value = newsData.count ?? 0;
  }
  isDownloadProcess.value = false;
}

const pageCount = computed(() => {
  try {
    return Math.ceil(totalCount.value / limit);
  } catch (error) {
    console.error(error);
  }

  return 0;
});

const imageUrl = (newItem) => {
  try {
    if (!newItem || !newItem.image) {
      return null;
    }

    return `${store.api_url}${newItem.image}`;
  } catch (error) {
    console.error(error);
    return null;
  }
}

</script>

<template>
  <div class="banner-2 space-y-10 pb-20" id="work">
    <h3 class="heading3 my-5">Новости</h3>
    <Loader v-if="isDownloadProcess" />
    <template v-else>
      <div class="card" v-for="(newItem, code) in news" :key="code">
        <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2">
          <h4 class="project-title item">{{newItem.title}}</h4>
          <p class="font-work_sans pr-12" v-html="newItem.description"></p>
          <a class="text-sky-800 font-bold text-2xl tracking-wider" :href="getNewUrl(newItem.id)">
            Подробнее
          </a>
        </div>
        <div v-if="imageUrl(newItem)" class="card-image bg-green-100">
          <img
              class="object-cover w-full h-72 md:h-96"
              :src=imageUrl(newItem)
          />
        </div>
      </div>
    </template>
    <Pagination v-if="pageCount > 1" :currentPage="currentPage" :pageCount="pageCount" :visiblePagesCount="5"
                @changePage="changePage"
    ></Pagination>
  </div>
</template>

<style scoped>

</style>