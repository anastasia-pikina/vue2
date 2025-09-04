<script setup>

import { computed } from 'vue'

const props = defineProps([
  'pageNumber',
  'pageNumberView',
])

const pageLink = computed(() => {
  return {
    name: 'news',
    params: {
      page: props.pageNumber,
    },
  }
});
</script>

<template>
<span>
		<router-link @click="$emit('changePage', pageNumber)" :to="pageLink">
			{{ pageNumberView }}
		</router-link>
	</span>
</template>

<style scoped>
  span.page a {
    padding: 0 .75em;
    border-right: 1px solid #cdcdcd;
  }

  span.page:last-child a {
    border: none;
  }

  span.page.page-current a
  {
    font-weight: bold;
    cursor: default;
  }

  a.disabled
  {
    color: #5c6470;
    cursor: default;
  }
</style>