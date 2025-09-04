<script setup>
import { ref, onMounted, computed } from 'vue'
import {useMainStore} from '../stores/mainStore';
const store = useMainStore();
const news = ref([]);
const maxNewIndex = ref(0);
const currentNewIndex = ref(0);
const limit = 10;

onMounted(async () => {
  try {
    const newsData = await store.getData(`/news?limit=${limit}`);
    if (newsData) {
      news.value = newsData.list ?? [];
      currentNew.value = news.value.find(x=>x!==undefined);
      maxNewIndex.value = news.value.length - 1;
    }
  } catch (error) {
    console.error(error);
  }
});

const currentNew = computed(() => {
  return news.value[currentNewIndex.value];
})


const getNewUrl = (NewId) => {
  return ['/news', NewId].join('/') + '/';
}

const nextNew = () => {
  if (maxNewIndex.value > currentNewIndex.value) {
    currentNewIndex.value++;
  }
}

const prevNew = () => {
  if (currentNewIndex.value > 0) {
    currentNewIndex.value--;
  }
}

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

const isMaxNew = computed(() => {
  return currentNewIndex.value === maxNewIndex.value;
});

const isMinNew = computed(() => {
  return currentNewIndex.value === 0;
});
</script>

<template>
  <div class="banner-1" id="testimonial" v-if="news.length > 0">
    <div class="space-y-4 grid place-items-center mt-5">
      <h3 class="heading3">Новости</h3>
      <p class="font-work_sans text-gray-600 font-semibold leading-relaxed">
        Последние новости компании
      </p>
    </div>
    <div class="max-w-6xl mx-auto px-8 py-16">
      <div class="relative">
        <div class="relative lg:flex rounded-lg shadow-2xl overflow-hidden">
          <div
              class="h-56 lg:h-auto lg:w-5/12 relative flex items-center justify-center"
          >
            <img v-if="imageUrl(currentNew)"
                class="absolute h-full w-full object-cover"
                :src=imageUrl(currentNew)
                alt=""
            />
            <div class="absolute inset-0 bg-gray-700 opacity-75"></div>
            <div
                class="relative main_news"
                width="200"
                height="120"
            >
              {{currentNew.title}}
            </div>
          </div>
          <div class="relative lg:w-7/12 bg-white">
            <svg
                class="absolute h-full text-white w-24 -ml-12"
                fill="currentColor"
                viewBox="0 0 100 100"
                preserveAspectRatio="none"
            >
              <polygon points="50,0 100,0 50,100 0,100" />
            </svg>
            <div
                class="relative py-12 lg:py-24 px-8 lg:px-16 text-gray-700 leading-relaxed"
            >
              <p v-if="currentNew.description">{{currentNew.description}}</p>
              <p class="mt-6">
                <a
                    :href="getNewUrl(currentNew.id)"
                    class="font-medium text-indigo-600 hover:text-indigo-900"
                >
                  &rarr; Подробнее</a>
              </p>
            </div>
          </div>
        </div>
        <div class="absolute inset-y-0 left-0 lg:flex lg:items-center">
          <button
              :class="{'is-disabled' : isMinNew}"
              @click="prevNew()"
              class="mt-24 lg:mt-0 -ml-6 h-12 w-12 rounded-full bg-white p-3 shadow-lg"
          >
            <svg
                class="h-full w-full text-indigo-900"
                fill="currentColor"
                viewBox="0 0 24 24"
            >
              <path
                  d="M5.41 11H21a1 1 0 0 1 0 2H5.41l5.3 5.3a1 1 0 0 1-1.42 1.4l-7-7a1 1 0 0 1 0-1.4l7-7a1 1 0 0 1 1.42 1.4L5.4 11z"
              />
            </svg>
          </button>
        </div>
        <div class="absolute inset-y-0 right-0 lg:flex lg:items-center">
          <button
              :class="{'is-disabled' : isMaxNew}"
              @click="nextNew()"
              class="mt-24 lg:mt-0 -mr-6 h-12 w-12 rounded-full bg-white p-3 shadow-lg"
          >
            <svg
                class="h-full w-full text-indigo-900"
                fill="currentColor"
                viewBox="0 0 24 24"
            >
              <path
                  d="M18.59 13H3a1 1 0 0 1 0-2h15.59l-5.3-5.3a1 1 0 1 1 1.42-1.4l7 7a1 1 0 0 1 0 1.4l-7 7a1 1 0 0 1-1.42-1.4l5.3-5.3z"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main_news {
  color: #fff;
  font-weight: 600;
  font-size: 42px;
}
</style>