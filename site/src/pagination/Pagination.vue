<script setup>
import { computed } from 'vue'

import PaginationPage from "./PaginationPage.vue";

const props = defineProps([
  'currentPage',
  'pageCount',
  'visiblePagesCount'
])

const paginationPages = computed(() => {

  let resultPages = {};

  try {
    const pageCount = props.pageCount;
    resultPages[1] = 1;
    if (pageCount === 1) {
      return resultPages;
    }

    const currentPage = props.currentPage;
    let visiblePagesCount = props.visiblePagesCount;
    const visiblePagesThreshold = (visiblePagesCount - 1) / 2;

    visiblePagesCount--;

    if (currentPage <= visiblePagesThreshold + 1 || pageCount <= props.visiblePagesCount + 1) {
      let lastPage = 1;

      while (visiblePagesCount >= 1) {
        resultPages[++lastPage] = lastPage;
        visiblePagesCount--;
        if (Object.keys(resultPages).length === pageCount) {
          return resultPages;
        }
      }

      resultPages[++lastPage] = '...';
      resultPages[pageCount] = pageCount;
      return resultPages;
    }

    if (currentPage >= pageCount - visiblePagesThreshold + 1) {
      let lastPage = pageCount - visiblePagesCount -1;
      resultPages[lastPage] = '...';

      while (visiblePagesCount >= 1) {
        resultPages[++lastPage] = lastPage;
        visiblePagesCount--;
      }

      resultPages[pageCount] = pageCount;
      return resultPages;
    }

    let lastPage = currentPage - visiblePagesThreshold;
    if (pageCount - currentPage === 2) {
      lastPage--;
    }
    resultPages[lastPage] = lastPage === 2 ? lastPage : '...';

    while (visiblePagesCount > 1) {
      resultPages[++lastPage] = lastPage;
      visiblePagesCount--;
    }

    resultPages[++lastPage] = (lastPage === pageCount - 1) ? lastPage : '...';
    resultPages[pageCount] = pageCount;
  } catch (e) {
    console.error(e.message);
  }

  return resultPages;
});

const pagePrevLink = computed(() => {
    try {
      let prevPage = props.currentPage - 1;
      if (prevPage < 1) {
        prevPage = 1;
      }

      return {
        name: 'news',
        params: {
          page: prevPage,
        },
      }
    } catch (e) {
      console.error(e.message);
    }

    return {};
});

const pageNextLink = computed(() => {
  try {
    let currentPage = props.currentPage;
    if (currentPage === 0) {
      currentPage = 1;
    }

    let nextPage = currentPage + 1;
    if (nextPage > props.pageCount) {
      nextPage = props.pageCount;
    }

    return {
      name: 'news',
      params: {
        page: nextPage,
      },
    }
  } catch (e) {
    console.error(e.message);
  }

  return {};
});

</script>

<template>
  <router-link :to="pagePrevLink"
               @click="$emit('changePage', currentPage - 1)"
               class="pages-nav">
    предыдущая
  </router-link>
  <PaginationPage
      v-for="(pageNumberView, pageNumber) in paginationPages" :key="pageNumber"
      :class="{
						'page-current':
						pageNumber === currentPage
					}"
      :pageNumber="pageNumber"
      :pageNumberView="pageNumberView"
      class="page"
      @changePage="$emit('changePage', pageNumber)"
  />
  <router-link :to="pageNextLink"
               class="pages-nav"
               @click="$emit('changePage', Number(currentPage) + 1)"
  >
    следующая
  </router-link>
</template>

<style scoped>

</style>