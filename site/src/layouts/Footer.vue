<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios';
const isDownload = ref(false)
const isIconsDownload = ref(false)
const contentFooter = ref();
const icons = ref();
const contentEmail = ref();

onMounted(() => {
  axios.get('http://localhost:4000/blocks/footer')
      .then(response => {
        contentFooter.value = response.data.find(x=>x!==undefined);
        console.log(contentFooter.value.content)

        axios.get('http://localhost:4000/blocks/email')
            .then(response => {
              contentEmail.value = response.data.find(x=>x!==undefined);
              isDownload.value = true;
             // console.log(contentEmail.value)
            })
            .catch(error => {
              console.log(error);
            });
      })
      .catch(error => {
        isDownload.value = true;
        console.log(error);
      });

  axios.get('http://localhost:4000/blocks/footer_icons')
      .then(response => {
        icons.value = response.data;
        isIconsDownload.value = true;
      })
      .catch(error => {
        isIconsDownload.value = true;
        console.log(error);
      });
})

</script>

<template>
    <!-- Footer Start -->
    <footer class="bg-theme-dark-blue">
      <div class="banner-1 flex h-full items-center">
        <div class="w-7/12 p-12" v-if="isDownload">
          <h2 class="text-gray-700 md:text-6xl text-2xl font-Eczar mb-5 font-bold">
            {{contentFooter.content}}
          </h2>
          <a v-if="contentEmail" class="underline text-2xl text-blue-600 font-work_sans"
          >{{contentEmail.content}}</a
          >
        </div>
        <div class="w-5/12 pr-28" v-if="isIconsDownload">
          <div class="flex flex-wrap justify-end gap-2">
            <a :href="icon.link" target="_blank" v-for="(icon, code) in icons" :key="code"
                class="bg-gray-700 p-2 font-semibold text-white inline-flex items-center space-x-2 rounded"
            ><span v-html="icon.title"></span>
            </a>
          </div>
        </div>
      </div>
    </footer>
    <!-- Footer End -->
</template>
