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

</style>