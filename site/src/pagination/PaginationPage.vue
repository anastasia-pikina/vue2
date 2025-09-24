<script setup>

import { computed } from 'vue'
import {useMainStore} from "../stores/mainStore";

const props = defineProps([
  'pageNumber',
  'pageNumberView',
]);

const store = useMainStore();

const pageLink = computed(() => {
  return store.getNewListUrl(props.pageNumber);
});
</script>

<template>
<span>
		<router-link :to="pageLink">
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