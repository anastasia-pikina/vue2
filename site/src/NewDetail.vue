<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useMainStore } from './stores/mainStore';
import { computed } from 'vue';
import Loader from './sections/Loader.vue';
const store = useMainStore();
const newDetail = ref(null);
const isDownload = ref(false);

onMounted(async () => {
  try {
    const route = useRoute()
    const id = route.params.id;
    newDetail.value = await store.getData(`/news/${id}`);
    isDownload.value = true;
  } catch (error) {
    return null;
  }
})

const imageUrl = computed(() => {
  try {
    if (!newDetail.value || !newDetail.value.image) {
      return null;
    }

    return `${store.api_url}${newDetail.value.image}`;
  } catch (error) {
    console.error(error);
    return null;
  }
})

</script>
<template>
  <div class="banner-2 space-y-10 pb-20" id="work">
    <h3 class="heading3 my-5">Новости</h3>
    <Loader v-if="!isDownload" />
    <div v-else class="card">
      <template v-if="newDetail">
        <div class="space-y-5 py-8 px-8 md:py-16 md:px-20 md:w-1/2">
          <h4 class="project-title item">{{newDetail.title}}</h4>
          <p class="font-work_sans pr-12" v-html="newDetail.content"></p>
        </div>
        <div v-if="imageUrl" class="card-image bg-green-100">
          <img
              class="object-cover w-full h-72 md:h-96"
              :src=imageUrl
          />
        </div>
      </template>
      <div v-else class="new-not-found">
        <p class="font-work_sans pr-12">Новость не найдена.</p>
      </div>
    </div>
  </div>
</template>
<style scoped>
  .new-not-found {
    padding: 15px;
  }
</style>