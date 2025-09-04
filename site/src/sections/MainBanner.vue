<script setup>
import {ref, onMounted, computed} from 'vue';
import {useMainStore} from '../stores/mainStore';
const store = useMainStore();
const contentBanner = ref(null);

onMounted(async () => {
  try {
    const contentBannerData = await store.getData('/blocks/main');
    if (contentBannerData) {
      contentBanner.value = contentBannerData.find(x => x !== undefined);
    }
  } catch (error) {
    console.error(error);
  }
})

const imageUrl = computed(() => {
  try {
    if (!contentBanner.value || !contentBanner.value.image) {
      return null;
    }

    return `${store.api_url}${contentBanner.value.image}`;
  } catch (error) {
    console.error(error);
    return null;
  }
})
</script>

<template>
  <div class="banner-2 md:flex items-center md:justify-evenly" id="about" v-if="contentBanner">
    <img v-if="imageUrl"
        class="md:h-[500px] h-[400px]"
        :src=imageUrl
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