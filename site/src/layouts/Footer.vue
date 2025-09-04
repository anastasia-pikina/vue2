<script setup>
import { ref, onMounted } from 'vue'
import { useMainStore } from '../stores/mainStore'
const store = useMainStore();
const contentFooter = ref(null);
const contentEmail = ref(null);
const icons = ref([]);

onMounted(async () => {
  const contentFooterData = await store.getData('/blocks/footer');
  if (contentFooterData) {
    contentFooter.value = contentFooterData.find(x => x !== undefined);
  }

  const contentEmailData = await store.getData('/blocks/email');
  if (contentEmailData) {
    contentEmail.value = contentEmailData.find(x => x !== undefined);
  }

  icons.value = (await store.getData('/blocks/footer_icons')) ?? [];
});

</script>

<template>
    <!-- Footer Start -->
    <footer class="bg-theme-dark-blue footer">
      <div class="banner-1 flex h-full items-center">
        <div class="w-7/12 p-12">
          <h2 v-if="contentFooter" class="text-gray-700 md:text-6xl text-2xl font-Eczar mb-5 font-bold">
            {{contentFooter.content}}
          </h2>
          <a v-if="contentEmail" class="underline text-2xl text-blue-600 font-work_sans"
          >{{contentEmail.content}}</a
          >
        </div>
        <div class="w-5/12 pr-28" v-if="icons.length > 0">
          <div class="flex flex-wrap justify-end gap-2">
            <a :href="icon.link" target="_blank" v-for="(icon, code) in icons" :key="code"
                class="bg-gray-700 p-2 font-semibold text-white inline-flex items-center space-x-2 rounded"
            ><span class="icon-svg" v-html="icon.title"></span>
            </a>
          </div>
        </div>
      </div>
    </footer>
    <!-- Footer End -->
</template>
<style scoped>
.footer .icon-svg {
  width: 24px;
}
</style>
